/*
 * Main menu view, ported from internal/ui/menu_view.go.
 *
 * Renders the three menu entries (About / Projects / Contact). The selected
 * entry is shown with a "> " prefix, bold purple text, and its description
 * underneath. Every entry is clickable/tappable and activates its screen.
 */

import { MENU_ENTRIES } from '../data'
import { Line } from './Line'
import { type ContentLine } from '../theme'

interface MenuViewProps {
  selected: number
  onActivateEntry: (index: number) => void
  onHighlightEntry: (index: number) => void
}

export function MenuView({
  selected,
  onActivateEntry,
  onHighlightEntry,
}: MenuViewProps) {
  const blankLine: ContentLine = { text: '', style: 'body' }
  const breakLine: ContentLine = { text: '\n', style: 'body' }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 0 }}>
      <Line line={blankLine} />
      <Line line={breakLine} />

      {MENU_ENTRIES.map((menu, i) => {
        const isSelected = i === selected

        return (
          <div
            key={i}
            onClick={(e) => {
              e.stopPropagation()
              onActivateEntry(i)
            }}
            onMouseEnter={() => onHighlightEntry(i)}
            style={{
              cursor: 'pointer',
              WebkitTapHighlightColor: 'transparent',
              marginBottom: '1.5rem',
              display: 'flex',
              flexDirection: 'column',
              alignSelf: 'flex-start',
              width: 'fit-content',
            }}
          >
            {isSelected && (
              <>
                <Line line={{ text: `> ${menu.title}`, style: 'selectedItem' as const }} />
                <Line line={{ text: menu.description, style: 'description' }} />
              </>
            )}
            {!isSelected && <Line line={{ text: `  ${menu.title}`, style: 'item' as const }} />}

            <Line line={breakLine} />
            <Line line={blankLine} />
          </div>
        )
      })}
    </div>
  )
}
