/*
 * Search input shown in the footer when searching is active, ported from the
 * searching branch of footerView() in internal/ui/layout.go.
 *
 * The placeholder mirrors the Bubble Tea textinput ("Search...  •  n/N to
 * navigate"). Enter submits/exits search; Escape cancels.
 */

interface SearchInputProps {
  value: string
  onChange: (value: string) => void
  onAction: (action: 'submit' | 'cancel') => void
}

export function SearchInput({ value, onChange, onAction }: SearchInputProps) {
  return (
    <input
      data-terminal-search="true"
      autoFocus
      placeholder="Search...  •  n/N to navigate"
      value={value}
      onChange={(e) => onChange(e.target.value)}
      onKeyDown={(e) => {
        if (e.key === 'Enter') {
          e.preventDefault()
          e.stopPropagation()
          onAction('submit')
        } else if (e.key === 'Escape') {
          e.preventDefault()
          e.stopPropagation()
          onAction('cancel')
        }
      }}
      style={{
        background: 'transparent',
        color: '#FFFFFF',
        border: 'none',
        outline: 'none',
        font: 'inherit',
        width: '100%',
        maxWidth: '30rem',
      }}
    />
  )
}
