/*
 * Terminal footer / help line, ported from footerView() in internal/ui/layout.go.
 *
 * Renders screen-specific key hints. Labels are muted, keys are bold blue.
 * When searching is active, the footer instead shows the search input.
 */

import { SearchInput } from './SearchInput'

type FooterItem = { label: string; key: string }

interface FooterProps {
  screen: 'menu' | 'about' | 'projects' | 'contact' | 'ssh'
  searching: boolean
  showFooter: boolean
  query: string
  onQueryChange: (value: string) => void
  onSearchAction: (action: 'submit' | 'cancel') => void
}

const HINTS: Record<string, FooterItem[]> = {
  menu: [
    { label: 'Navigate', key: '↑/↓ j/k/Wheel' },
    { label: 'Select', key: 'Click/Enter' },
    { label: 'Back', key: 'b/Esc' },
  ],
  about: [
    { label: 'Navigate', key: '↑/↓ j/k' },
    { label: 'Scroll', key: 'Ctrl+D/U' },
    { label: 'Back', key: 'b/Esc' },
    { label: 'Search', key: '/' },
  ],
  projects: [
    { label: 'Navigate', key: '↑/↓ j/k/Wheel' },
    { label: 'Open GitHub', key: 'Enter/Click' },
    { label: 'Back', key: 'b/Esc' },
    { label: 'Search', key: '/' },
  ],
  ssh: [
    { label: 'Navigate', key: '↑/↓ j/k/Wheel' },
    { label: 'Scroll', key: 'Ctrl+D/U' },
    { label: 'Back', key: 'b/Esc' },
  ],
  contact: [
    { label: 'Next', key: 'Tab' },
    { label: 'Previous', key: 'Shift+Tab' },
    { label: 'Back', key: 'b/Esc' },
    { label: 'Send', key: 'Enter' },
  ],
}

const GAP = '    '

export function Footer({
  screen,
  searching,
  showFooter,
  query,
  onQueryChange,
  onSearchAction,
}: FooterProps) {
  if (!showFooter) {
    return null
  }

  if (searching) {
    return (
      <div style={{ color: '#565F89' }}>
        <span style={{ color: '#2D7DFF', fontWeight: 'bold' }}>› </span>
        <SearchInput
          value={query}
          onChange={onQueryChange}
          onAction={onSearchAction}
        />
      </div>
    )
  }

  const items = [...HINTS[screen], { label: '? Help', key: '' }]

  return (
    <footer style={{ color: '#565F89', padding: '1rem 1.5rem', boxSizing: 'border-box' }}>
      {items.map((item, i) => (
        <span key={i} style={{ whiteSpace: 'nowrap' }}>
          <span style={{ color: '#A9B1D6' }}>{item.label}</span>{' '}
          {item.key ? (
            <span style={{ color: '#2D7DFF', fontWeight: 'bold' }}>{item.key}</span>
          ) : null}
          {i < items.length - 1 ? <span style={{ color: '#565F89' }}>{GAP}</span> : null}
        </span>
      ))}
    </footer>
  )
}
