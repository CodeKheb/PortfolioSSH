/*
 * About view, ported from internal/ui/about_view.go.
 *
 * Renders the identity, introduction, education, experiences, skillset and
 * certifications. In the terminal the block is centered when the terminal is
 * wider than 80 columns; here that is expressed as a centered 80ch column that
 * naturally becomes full width on narrow screens.
 */

import { useCallback, useEffect } from 'react'
import { ABOUT_LINES } from '../data'
import { Line } from './Line'
import { type ContentLine } from '../theme'

interface AboutViewProps {
  query: string
  registerSearchables: (getter: () => readonly ContentLine[]) => void
}

export function AboutView({ query, registerSearchables }: AboutViewProps) {
  const getSearchableLines = useCallback(() => ABOUT_LINES, [])

  useEffect(() => {
    registerSearchables(getSearchableLines)
  }, [registerSearchables, getSearchableLines])

  return (
    <div className="about-view" style={{ display: 'flex', flexDirection: 'column', gap: 0 }}>
      <div
        style={{
          maxWidth: '80ch',
          margin: '0 auto',
          width: '100%',
        }}
      >
        {ABOUT_LINES.map((line, i) => (
          <div key={i}>
            <Line line={line} query={query} className="term-line" />
          </div>
        ))}
      </div>
    </div>
  )
}
