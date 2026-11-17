/*
 * A small, terminal-styled "Back" affordance for touch / mouse navigation.
 * Rendered in the top right of the header on every non-menu screen so the
 * interface is usable without a keyboard (mirrors the "Back b/Esc" hint).
 */

interface BackButtonProps {
  onBack: () => void
}

export function BackButton({ onBack }: BackButtonProps) {
  return (
    <div
      onClick={onBack}
      onKeyDown={(e) => {
        if (e.key === 'Enter' || e.key === ' ') {
          e.preventDefault()
          onBack()
        }
      }}
      role="button"
      tabIndex={0}
      aria-label="Back"
      style={{
        color: '#1A1B26', // DARK
        fontSize: '0.9rem',
        fontWeight: 'bold',
        cursor: 'pointer',
        userSelect: 'none',
        WebkitTapHighlightColor: 'transparent',
        display: 'inline-flex',
        alignItems: 'center',
        gap: '0.25rem',
        padding: '0.15rem 0.6rem',
        border: '1px solid #1A1B26',
        flexShrink: 0,
      }}
    >
      ← Back
    </div>
  )
}
