/*
 * Terminal theme constants.
 *
 * These mirror the Lipgloss color palette defined in internal/ui/styles.go of the
 * Go SSH portfolio so the browser port looks identical to the terminal version.
 */

export const theme = {
  // Palette (exact hex values from styles.go)
  BLUE: '#2D7DFF',
  PURPLE: '#7D56F4',
  GREEN: '#73DACA',
  RED: '#F7768E',
  WHITE: '#FFFFFF',
  CREAM: '#FFFDF5',
  TEXT: '#C0CAF5',
  MUTED: '#A9B1D6',
  DIM: '#565F89',
  DARK: '#1A1B26',
  LAVENDER: '#BB9AF7',
} as const

// Terminal screen identifiers, mirroring the Screen iota in internal/ui/model.go
export type Screen = 'menu' | 'about' | 'projects' | 'readme' | 'contact'

export const SCREENS: Screen[] = ['menu', 'about', 'projects', 'readme', 'contact']

// Named line styles. Each maps to the equivalent Lipgloss style in styles.go.
export type LineStyle =
  | 'body'
  | 'muted'
  | 'name'
  | 'section'
  | 'subheading'
  | 'code'
  | 'bullet'
  | 'description'
  | 'descriptionAligned'
  | 'item'
  | 'selectedItem'
  | 'title'
  | 'contactLabel'
  | 'contactFocusedLabel'
  | 'contactError'
  | 'contactSuccess'

// One rendered content line, mirroring ContentLine in internal/ui/search.go.
// imageSrc/imageAlt are used for markdown images (screenshots etc.), which the
// terminal cannot show but the browser can.
export interface ContentLine {
  text: string
  style: LineStyle
  nowrap?: boolean
  imageSrc?: string
  imageAlt?: string
}

// A selectable list entry used by the menu and projects screens.
export interface MenuEntry {
  title: string
  description?: string
}

export interface Project {
  title: string
  technology: string
  description: string
  repoURL: string
  readme?: string
}
