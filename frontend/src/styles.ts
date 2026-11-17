/*
 * React style objects derived from the terminal theme.
 * These mirror the lipgloss styles in internal/ui/styles.go.
 */

import { theme, type LineStyle } from './theme'

// 1 terminal column ≈ 1 character width. Use `ch` units (exact width of the
// "0" glyph in the monospace font) so text columns line up exactly regardless
// of font metrics, mirroring lipgloss character-cell padding.
const CH = '1ch'
const LINE = 1.25

export const lineStyles: Record<LineStyle, React.CSSProperties> = {
  body: {
    color: theme.WHITE,
  },
  muted: {
    color: theme.MUTED,
  },
  name: {
    color: theme.BLUE,
    fontWeight: 'bold',
    marginBottom: LINE,
  },
  section: {
    color: theme.LAVENDER,
    fontWeight: 'bold',
    marginTop: LINE,
    marginBottom: LINE,
  },
  subheading: {
    color: theme.GREEN,
    fontWeight: 'bold',
  },
  code: {
    color: theme.MUTED,
    paddingLeft: CH,
  },
  bullet: {
    color: theme.RED,
  },
  description: {
    color: theme.TEXT,
    paddingLeft: '4ch',
  },
  descriptionAligned: {
    color: theme.TEXT,
    paddingLeft: '2ch',
  },
  item: {
    color: theme.WHITE,
    paddingLeft: '2ch',
  },
  selectedItem: {
    color: theme.PURPLE,
    fontWeight: 'bold',
    paddingLeft: '2ch',
  },
  title: {
    color: theme.CREAM,
    backgroundColor: theme.PURPLE,
    fontWeight: 'bold',
    padding: '0 4ch', // Padding(0, 4)
  },
  contactLabel: {
    color: theme.MUTED,
  },
  contactFocusedLabel: {
    color: theme.LAVENDER,
    fontWeight: 'bold',
  },
  contactError: {
    color: theme.RED,
  },
  contactSuccess: {
    color: theme.GREEN,
  },
}
