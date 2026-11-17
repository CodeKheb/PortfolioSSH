/*
 * PortfolioSSH view — how to install an SSH client, connect to the portfolio
 * over SSH and save the server's host key.
 *
 * Styled like a terminal.shop session: a numbered script with dim comments and
 * click-to-copy commands.
 */

import { useEffect, useRef, useState } from 'react'

export const HOST = 'kherbin.getclingy.download'
export const HOST_KEY =
  'ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIGgglhVh0jaTSuosSZUq5MYtMah97v16Ac6vjmWtn3W8'

type RowKind = 'comment' | 'command' | 'output' | 'link' | 'blank'

interface Row {
  text: string
  kind: RowKind
}

const ROWS: readonly Row[] = [
  { text: '# install openssh (skip if ssh is already installed)', kind: 'comment' },
  { text: '# windows (powershell)', kind: 'comment' },
  { text: 'Add-WindowsCapability -Online -Name OpenSSH.Client~~~~0.0.1.0', kind: 'command' },
  { text: '# powershell not working? install it from the microsoft docs', kind: 'comment' },
  {
    text: 'https://learn.microsoft.com/en-us/windows-server/administration/openssh/openssh_install_firstuse',
    kind: 'link',
  },
  { text: '# macos', kind: 'comment' },
  { text: 'brew install openssh', kind: 'command' },
  { text: '# linux', kind: 'comment' },
  { text: 'sudo apt install openssh-client', kind: 'command' },
  { text: 'sudo dnf install openssh-clients', kind: 'command' },
  { text: 'sudo pacman -S openssh', kind: 'command' },
  { text: '', kind: 'blank' },

  { text: '# connect to the portfolio, no password needed', kind: 'comment' },
  { text: `ssh ${HOST}`, kind: 'command' },
  { text: '', kind: 'blank' },

  { text: '# first connect asks to trust the host, type yes to save it', kind: 'comment' },
  { text: `ssh-keyscan -t ed25519 ${HOST}`, kind: 'command' },
  { text: '', kind: 'blank' },

  { text: '# or save the host key to known_hosts directly', kind: 'comment' },
  { text: `ssh-keyscan -t ed25519 ${HOST} >> ~/.ssh/known_hosts`, kind: 'command' },
  { text: '', kind: 'blank' },

  { text: '# the expected host key looks like this', kind: 'comment' },
  { text: `${HOST} ${HOST_KEY}`, kind: 'output' },
  { text: '', kind: 'blank' },

  { text: '# forgot the key? remove the old entry and reconnect', kind: 'comment' },
  { text: `ssh-keygen -R ${HOST}`, kind: 'command' },
  { text: '', kind: 'blank' },

  { text: '# click any command to copy it', kind: 'comment' },
]

const textStyle: Record<RowKind, React.CSSProperties> = {
  comment: { color: '#565F89' }, // DIM
  command: { color: '#FFFFFF', fontWeight: 'bold' }, // WHITE
  output: { color: '#A9B1D6' }, // MUTED
  link: { color: '#2D7DFF', fontWeight: 'bold' }, // BLUE
  blank: {},
}

function copyText(text: string): Promise<void> {
  if (navigator.clipboard?.writeText) {
    return navigator.clipboard.writeText(text).catch(() => fallbackCopy(text))
  }
  fallbackCopy(text)
  return Promise.resolve()
}

function fallbackCopy(text: string) {
  const field = document.createElement('textarea')
  field.value = text
  field.style.position = 'fixed'
  field.style.opacity = '0'
  document.body.appendChild(field)
  field.select()
  try {
    document.execCommand('copy')
  } catch {
    // Clipboard access was denied; nothing else to try.
  }
  document.body.removeChild(field)
}

export function SshView() {
  const [copied, setCopied] = useState<number | null>(null)
  const copiedTimer = useRef<number | null>(null)

  useEffect(() => {
    return () => {
      if (copiedTimer.current !== null) window.clearTimeout(copiedTimer.current)
    }
  }, [])

  const handleCopy = (index: number, text: string) => {
    void copyText(text)
    setCopied(index)
    if (copiedTimer.current !== null) window.clearTimeout(copiedTimer.current)
    copiedTimer.current = window.setTimeout(() => setCopied(null), 1400)
  }

  return (
    <div
      className="ssh-view"
      style={{
        display: 'flex',
        flexDirection: 'column',
        maxWidth: '110ch',
        width: '100%',
        margin: '0.5rem auto 0',
        gap: '0.35rem',
      }}
    >
      {ROWS.map((row, i) => {
        const copyable = row.kind === 'command' || row.kind === 'output'
        const isLink = row.kind === 'link'

        return (
          <div
            key={i}
            style={{
              display: 'flex',
              alignItems: 'flex-start',
              gap: '2ch',
            }}
          >
            <span
              style={{
                color: '#565F89', // DIM
                minWidth: '3ch',
                textAlign: 'right',
                flexShrink: 0,
                userSelect: 'none',
              }}
            >
              {i + 1}
            </span>

            <span
              onClick={
                copyable
                  ? () => handleCopy(i, row.text)
                  : isLink
                    ? () => window.open(row.text, '_blank', 'noopener,noreferrer')
                    : undefined
              }
              title={copyable ? 'Click to copy' : isLink ? 'Open in a new tab' : undefined}
              style={{
                ...textStyle[row.kind],
                whiteSpace: 'pre-wrap',
                overflowWrap: 'break-word',
                minWidth: 0,
                cursor: copyable || isLink ? 'pointer' : undefined,
                textDecoration: isLink ? 'underline' : undefined,
                WebkitTapHighlightColor: 'transparent',
              }}
            >
              {row.text}
              {copied === i && (
                <span
                  style={{
                    color: '#73DACA', // GREEN
                    fontWeight: 'normal',
                    marginLeft: '1ch',
                    userSelect: 'none',
                  }}
                >
                  [copied]
                </span>
              )}
            </span>
          </div>
        )
      })}
    </div>
  )
}
