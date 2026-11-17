/*
 * Renders a single terminal content line (ContentLine) with optional live
 * search highlighting, mirroring renderLines() / highlightSearch() in
 * internal/ui/search.go.
 */

import { type ContentLine } from '../theme'
import { lineStyles } from '../styles'

export interface LineProps {
  line: ContentLine
  query?: string
  dataset?: Record<string, string>
  className?: string
}

function highlightQuery(text: string, query?: string): Array<React.ReactNode> {  if (!query) {
    return [text]
  }

  const trimmed = query.trim()
  if (trimmed === '') {
    return [text]
  }

  const lowerText = text.toLowerCase()
  const lowerQuery = trimmed.toLowerCase()

  const nodes: Array<React.ReactNode> = []
  let start = 0
  let index: number

  while ((index = lowerText.indexOf(lowerQuery, start)) !== -1) {
    if (index > start) {
      nodes.push(text.slice(start, index))
    }
    nodes.push(
      <mark key={start} style={styles.searchHighlight}>
        {text.slice(index, index + trimmed.length)}
      </mark>,
    )
    start = index + trimmed.length
  }

  if (start < text.length) {
    nodes.push(text.slice(start))
  }

  return nodes
}

const styles = {
  searchHighlight: {
    backgroundColor: '#2D7DFF', // theme.BLUE
    color: '#FFFFFF', // theme.WHITE
    fontWeight: 'bold' as const,
  } as React.CSSProperties,
}

export function Line({ line, query, dataset, className }: LineProps) {
  const base = lineStyles[line.style]

  // Markdown images (screenshots, etc.) cannot be shown in a terminal, but the
  // browser can render them inline.
  if (line.imageSrc) {
    return (
      <div
        className={className}
        {...dataset}
        style={{
          ...base,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'flex-start',
          gap: '0.25rem',
        }}
      >
        <img
          src={line.imageSrc}
          alt={line.imageAlt || ''}
          style={{
            maxWidth: '100%',
            height: 'auto',
            border: '1px solid #565F89',
            display: 'block',
            margin: 0,
          }}
          loading="lazy"
        />
        {line.imageAlt && (
          <div style={{ fontSize: '0.85rem', color: '#A9B1D6', marginTop: '0.25rem' }}>
            {line.imageAlt}
          </div>
        )}
      </div>
    )
  }

  return (
    <div
      className={className}
      {...dataset}
      style={{
        ...base,
        whiteSpace: 'pre-wrap',
      }}
    >
      {highlightQuery(line.text, query)}
    </div>
  )
}
