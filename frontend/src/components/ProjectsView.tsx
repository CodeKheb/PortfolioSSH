/*
 * Projects view — the selected project is expanded with its technology and
 * description; clicking a project opens its GitHub repository in a new tab.
 */

import { PROJECTS } from '../data'
import { Line } from './Line'
import { type ContentLine } from '../theme'

interface ProjectsViewProps {
  selected: number
  query: string
  onSelect: (index: number) => void
}

export function ProjectsView({ selected, query, onSelect }: ProjectsViewProps) {
  const handleProjectClick = (index: number) => {
    const project = PROJECTS[index]
    if (project.repoURL) {
      window.open(project.repoURL, '_blank', 'noopener,noreferrer')
    }
  }

  const titleLine = (index: number): ContentLine => {
    const project = PROJECTS[index]
    return index === selected
      ? { text: `> ${project.title}`, style: 'selectedItem' }
      : { text: `  ${project.title}`, style: 'item' }
  }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 0 }}>
      <Line line={{ text: '\n', style: 'body' }} />

      {PROJECTS.map((project, i) => {
        const isSelected = i === selected

        return (
          <div
            key={project.title}
            onMouseEnter={() => onSelect(i)}
            style={{
              marginBottom: '1.5rem',
              alignSelf: 'flex-start',
              width: 'fit-content',
              maxWidth: '100%',
            }}
          >
            <div
              data-proj={i}
              onClick={(e) => {
                e.stopPropagation()
                handleProjectClick(i)
              }}
              style={{
                cursor: 'pointer',
                WebkitTapHighlightColor: 'transparent',
                width: 'fit-content',
              }}
            >
              <Line
                line={titleLine(i)}
                query={query}
                className="term-line"
              />
            </div>

            {isSelected && (
              <>
                <Line
                  line={{ text: project.technology, style: 'description' }}
                  query={query}
                  className="term-line"
                />
                <Line
                  line={{ text: ' ', style: 'description' }}
                  className="term-line"
                />
                <Line
                  line={{ text: project.description, style: 'description' }}
                  query={query}
                  className="term-line"
                />
              </>
            )}
          </div>
        )
      })}

      <div style={{ marginTop: '0.5rem', fontSize: '0.9rem', color: '#565F89' }}>
        <span style={{ color: '#2D7DFF', fontWeight: 'bold' }}>Click/Enter</span>{' to open GitHub repo'}
      </div>
    </div>
  )
}
