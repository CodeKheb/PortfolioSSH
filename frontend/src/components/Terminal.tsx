/*
 * Terminal shell: the single-page "continuous terminal session" that hosts the
 * ported SSH portfolio screens.
 *
 * Owns the navigation state (screen, selected index), the search state, the
 * scrollable content areas (one per view, for independent mouse scrolling),
 * and the global keyboard handler. Individual views are pure components.
 */

import { useCallback, useEffect, useRef, useState } from 'react'
import { MENU_ENTRIES, PROJECTS } from '../data'
import { Header } from './Header'
import { Footer } from './Footer'
import { MenuView } from './MenuView'
import { AboutView } from './AboutView'
import { ProjectsView } from './ProjectsView'
import { ContactView } from './ContactView'
import { type ContentLine } from '../theme'

// One text line per arrow-key press.
const SCROLL_STEP = 36
// Wheel travel (in px) needed before the selection moves to the next entry.
const WHEEL_STEP = 40
// Minimum time between two wheel-driven selection moves (ms).
const WHEEL_LOCK = 160

type ViewScreen = 'menu' | 'about' | 'projects' | 'contact'

export function Terminal() {
  const [screen, setScreen] = useState<ViewScreen>('menu')
  const [selected, setSelected] = useState(0)
  const [showFooter, setShowFooter] = useState(true)
  const [searching, setSearching] = useState(false)
  const [query, setQuery] = useState('')

  // One ref per view so each view has its own independently scrollable region
  // for mouse wheel scrolling.
  const contentRefs = useRef<Record<ViewScreen, HTMLDivElement | null>>({
    menu: null,
    about: null,
    projects: null,
    contact: null,
  })
  // Getter for the currently mounted content view's searchable lines.
  const searchableGetterRef = useRef<(() => readonly ContentLine[]) | null>(null)
  // Match index for n/N navigation, persisted across renders.
  const searchMatchIndexRef = useRef(0)
  // Accumulated wheel travel / cooldown for wheel navigation.
  const wheelAccumRef = useRef(0)
  const wheelLockRef = useRef(0)

  const registerSearchables = useCallback(
    (getter: () => readonly ContentLine[]) => {
      searchableGetterRef.current = getter
    },
    [],
  )

  const activeContentRef = useCallback(() => {
    return contentRefs.current[screen]
  }, [screen])

  const gotoTop = useCallback(() => {
    const el = activeContentRef()
    if (el) {
      el.scrollTop = 0
    }
  }, [activeContentRef])

  const scrollByLines = useCallback(
    (lines: number) => {
      const el = activeContentRef()
      if (!el) return
      el.scrollBy({ top: lines * SCROLL_STEP, behavior: 'smooth' })
    },
    [activeContentRef],
  )

  const scrollHalfPage = useCallback(
    (dir: 'down' | 'up') => {
      const el = activeContentRef()
      if (!el) return
      const step = Math.round(el.clientHeight / 2) || SCROLL_STEP * 4
      el.scrollBy({ top: (dir === 'down' ? 1 : -1) * step, behavior: 'smooth' })
    },
    [activeContentRef],
  )

  const clampSelected = useCallback(
    (index: number, max: number) => Math.max(0, Math.min(max, index)),
    [],
  )

  const goBack = useCallback(() => {
    if (screen !== 'menu') {
      setScreen('menu')
      setSelected(0)
      gotoTop()
    }
  }, [screen, gotoTop])

  const activateMenuEntry = useCallback(
    (index: number) => {
      switch (index) {
        case 0:
          setSelected(0)
          setScreen('about')
          break
        case 1:
          setSelected(0)
          setScreen('projects')
          break
        case 2:
          setSelected(0)
          setScreen('contact')
          break
      }
    },
    [],
  )

  const selectEntry = useCallback(() => {
    if (screen === 'menu') {
      activateMenuEntry(selected)
      return
    }
    if (screen === 'projects') {
      const project = PROJECTS[selected]
      if (project?.repoURL) {
        window.open(project.repoURL, '_blank', 'noopener,noreferrer')
      }
    }
  }, [screen, selected, activateMenuEntry])

  // Build the full list of searchable lines for the current screen.
  // For the projects screen, one combined line per project is used so that a
  // match index maps 1:1 to a project index and a description match selects
  // (and reveals) the project it belongs to.
  const getSearchableLines = useCallback((): readonly ContentLine[] => {
    if (screen === 'projects') {
      const lines: ContentLine[] = []
      PROJECTS.forEach((project) => {
        lines.push({ text: `${project.title} ${project.technology} ${project.description}`, style: 'item' })
      })
      return lines
    }

    const getter = searchableGetterRef.current
    if (!getter) return []
    return getter()
  }, [screen])

  const findMatches = useCallback(
    (lines: readonly ContentLine[], q: string): number[] => {
      const qLower = q.trim().toLowerCase()
      if (qLower === '') return []
      const matches: number[] = []
      lines.forEach((line, i) => {
        if (line.text.toLowerCase().includes(qLower)) matches.push(i)
      })
      return matches
    },
    [],
  )

  // Projects re-render after the selection changes (the expanded description
  // shifts every following entry), so the jump has to run after that layout.
  useEffect(() => {
    if (screen !== 'projects') return
    const el = contentRefs.current.projects
    if (!el) return
    const raf = requestAnimationFrame(() => {
      const target = el.querySelector(`[data-proj="${selected}"]`)
      target?.scrollIntoView({ block: 'nearest', behavior: 'smooth' })
    })
    return () => cancelAnimationFrame(raf)
  }, [screen, selected])

  const scrollToLine = useCallback(
    (target: number) => {
      const el = activeContentRef()
      if (!el) return
      if (screen === 'projects') {
        setSelected(target)
        return
      }
      requestAnimationFrame(() => {
        const lineEls = el.querySelectorAll('.term-line')
        lineEls[target]?.scrollIntoView({ block: 'center', behavior: 'smooth' })
      })
    },
    [screen, activeContentRef],
  )

  const jumpToMatch = useCallback(
    (dir: 1 | -1) => {
      const lines = getSearchableLines()
      const matches = findMatches(lines, query)
      if (matches.length === 0) {
        searchMatchIndexRef.current = 0
        return
      }

      searchMatchIndexRef.current =
        (searchMatchIndexRef.current + dir + matches.length) % matches.length

      const target = matches[searchMatchIndexRef.current]
      scrollToLine(target)
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [query, getSearchableLines, findMatches, scrollToLine, screen],
  )

  const handleKeyDown = useCallback(
    (e: KeyboardEvent) => {
      // Escape always works, even from a focused input (back / exit search).
      if (e.key === 'Escape') {
        e.preventDefault()
        if (searching) {
          setSearching(false)
          setQuery('')
          return
        }
        goBack()
        return
      }

      // Enter while searching: exit search mode but KEEP the highlights
      // (mirrors the terminal, where Enter exits search without clearing).
      if (searching && e.key === 'Enter') {
        e.preventDefault()
        setSearching(false)
        return
      }

      const active = document.activeElement as HTMLElement | null
      const isInput =
        active?.tagName === 'INPUT' || active?.tagName === 'TEXTAREA'

      // Let focused inputs handle their own keys (the contact view handles
      // Tab/Shift+Tab/Enter for field cycling).
      if (isInput) return

      if (e.ctrlKey && (e.key === 'd' || e.key === 'u')) {
        e.preventDefault()
        scrollHalfPage(e.key === 'd' ? 'down' : 'up')
        return
      }

      switch (e.key) {
        case 'ArrowUp':
        case 'k':
          e.preventDefault()
          if (screen === 'menu') {
            setSelected((s) => clampSelected(s - 1, MENU_ENTRIES.length - 1))
          } else if (screen === 'projects') {
            setSelected((s) => clampSelected(s - 1, PROJECTS.length - 1))
          } else if (screen === 'about') {
            scrollByLines(-1)
          }
          break

        case 'ArrowDown':
        case 'j':
          e.preventDefault()
          if (screen === 'menu') {
            setSelected((s) => clampSelected(s + 1, MENU_ENTRIES.length - 1))
          } else if (screen === 'projects') {
            setSelected((s) => clampSelected(s + 1, PROJECTS.length - 1))
          } else if (screen === 'about') {
            scrollByLines(1)
          }
          break

        case 'Enter':
          e.preventDefault()
          selectEntry()
          break

        case 'b':
          e.preventDefault()
          goBack()
          break

        case 'q':
          e.preventDefault()
          // Quit / disconnect is not available inside a browser tab.
          break

        case '?':
          // Ignore auto-repeat so the help footer doesn't toggle rapidly.
          if (e.repeat) break
          e.preventDefault()
          setShowFooter((f) => !f)
          break

        case '/':
          // Ignore auto-repeat so search isn't restarted repeatedly.
          if (e.repeat) break
          e.preventDefault()
          if (screen === 'about' || screen === 'projects') {
            setSearching(true)
            setQuery('')
          }
          break

        case 'n':
        case 'N':
          // Only jump between matches when NOT actively typing in the search
          // input. Press Enter first to confirm the search (keeps highlights),
          // then n/N cycle through matches.
          e.preventDefault()
          jumpToMatch(e.key === 'n' ? 1 : -1)
          break
      }
    },
    [
      searching,
      goBack,
      selectEntry,
      screen,
      scrollHalfPage,
      scrollByLines,
      clampSelected,
      jumpToMatch,
    ],
  )

  useEffect(() => {
    window.addEventListener('keydown', handleKeyDown)
    return () => window.removeEventListener('keydown', handleKeyDown)
  }, [handleKeyDown])

  // Wheel navigation: in the menu and projects screens the wheel moves the
  // selection (one entry per gesture) instead of scrolling the list, exactly
  // like pressing j/k.
  useEffect(() => {
    if (screen !== 'menu' && screen !== 'projects') return

    const onWheel = (e: WheelEvent) => {
      const el = contentRefs.current[screen]
      if (!el || !el.contains(e.target as Node)) return

      e.preventDefault()

      const now = performance.now()
      if (now < wheelLockRef.current) return

      const acc = wheelAccumRef.current
      if (acc !== 0 && Math.sign(acc) !== Math.sign(e.deltaY)) {
        wheelAccumRef.current = 0
      }
      wheelAccumRef.current += e.deltaY

      if (Math.abs(wheelAccumRef.current) < WHEEL_STEP) return

      const dir = wheelAccumRef.current > 0 ? 1 : -1
      wheelAccumRef.current = 0
      wheelLockRef.current = now + WHEEL_LOCK

      if (screen === 'menu') {
        setSelected((s) => clampSelected(s + dir, MENU_ENTRIES.length - 1))
      } else {
        setSelected((s) => clampSelected(s + dir, PROJECTS.length - 1))
      }
    }

    window.addEventListener('wheel', onWheel, { passive: false })
    return () => window.removeEventListener('wheel', onWheel)
  }, [screen, clampSelected])

  // Recompute matches and jump to the first one as the user types, mirroring
  // searchContent() / searchProjects() in internal/ui/search.go.
  useEffect(() => {
    if (!searching) return
    searchMatchIndexRef.current = 0
    if (query.trim() === '') {
      gotoTop()
      return
    }
    const lines = getSearchableLines()
    const matches = findMatches(lines, query)
    if (matches.length === 0) return
    const raf = requestAnimationFrame(() => scrollToLine(matches[0]))
    return () => cancelAnimationFrame(raf)
  }, [searching, query, screen, getSearchableLines, findMatches, scrollToLine, gotoTop])

  // Each view gets its own independently scrollable container so mouse wheel
  // scrolling stays within that view's content rather than scrolling the
  // whole page or another view's content. `minHeight: 0` is required for the
  // flex child to shrink and actually become scrollable.
  const contentStyle: React.CSSProperties = {
    flex: 1,
    minHeight: 0,
    overflowY: 'auto',
    overflowX: 'hidden',
    padding: '1rem 1.5rem',
    boxSizing: 'border-box',
  }

  return (
    <div
      className="terminal"
      style={{
        display: 'flex',
        flexDirection: 'column',
        height: '100svh',
        minHeight: '100vh',
        overflow: 'hidden',
      }}
    >
      <Header screen={screen} projectTitle={undefined} onBack={goBack} />

      {screen === 'menu' && (
        <div
          ref={(el) => {
            contentRefs.current.menu = el
          }}
          onClick={selectEntry}
          style={contentStyle}
        >
          <MenuView
            selected={selected}
            onActivateEntry={activateMenuEntry}
            onHighlightEntry={setSelected}
          />
        </div>
      )}
      {screen === 'about' && (
        <div
          ref={(el) => {
            contentRefs.current.about = el
          }}
          style={contentStyle}
        >
          <AboutView
            query={query}
            registerSearchables={registerSearchables}
          />
        </div>
      )}
      {screen === 'projects' && (
        <div
          ref={(el) => {
            contentRefs.current.projects = el
          }}
          onClick={selectEntry}
          style={contentStyle}
        >
          <ProjectsView selected={selected} query={query} onSelect={setSelected} />
        </div>
      )}
      {screen === 'contact' && (
        <div
          ref={(el) => {
            contentRefs.current.contact = el
          }}
          style={contentStyle}
        >
          <ContactView />
        </div>
      )}

      <Footer
        screen={screen}
        searching={searching}
        showFooter={showFooter}
        query={query}
        onQueryChange={(value) => setQuery(value)}
        onSearchAction={(action) => {
          if (action === 'cancel') {
            setSearching(false)
            setQuery('')
          } else {
            setSearching(false)
          }
        }}
      />
    </div>
  )
}
