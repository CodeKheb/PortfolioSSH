/*
 * Terminal header, ported from headerView() in internal/ui/layout.go.
 * Full-width purple bar with dark bold title text. On every non-menu screen the
 * Back control sits at the top right of the bar.
 */

import { headerTitle } from '../data'
import { BackButton } from './BackButton'
import { type Screen } from '../theme'

interface HeaderProps {
  screen: Screen
  projectTitle?: string
  onBack: () => void
}

export function Header({ screen, projectTitle, onBack }: HeaderProps) {
  return (
    <header
      style={{
        backgroundColor: '#7D56F4', // PURPLE
        color: '#1A1B26', // DARK
        fontWeight: 'bold',
        width: '100%',
        padding: '1rem 1.5rem', // Padding(1, 2)
        boxSizing: 'border-box',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        gap: '1rem',
      }}
    >
      <span
        style={{
          minWidth: 0,
          overflow: 'hidden',
          textOverflow: 'ellipsis',
          whiteSpace: 'nowrap',
        }}
      >
        {headerTitle(screen, projectTitle)}
      </span>

      {screen !== 'menu' && <BackButton onBack={onBack} />}
    </header>
  )
}
