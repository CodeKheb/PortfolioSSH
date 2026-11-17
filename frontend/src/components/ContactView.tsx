/*
 * Contact view, ported from internal/ui/contact_view.go.
 *
 * A terminal-style message form (Name / Email / Message / Send). Tab and
 * Shift+Tab cycle the fields, Enter moves to the next field (or sends when on
 * the send button). Validation errors render in red; a successful send renders
 * "Thank u for da rose" in green, mirroring the terminal.
 */

import { useEffect, useRef, useState } from 'react'

const FIELD_NAME = 0
const FIELD_EMAIL = 1
const FIELD_MESSAGE = 2
const FIELD_SEND = 3

const indicatorStyle = (focused: boolean): React.CSSProperties => ({
  color: focused ? '#BB9AF7' : '#565F89', // LAVENDER / DIM
  fontWeight: focused ? 'bold' : 'bold',
  userSelect: 'none',
  marginRight: '0.5rem',
} as const)

const labelStyle = (focused: boolean): React.CSSProperties => ({
  color: focused ? '#BB9AF7' : '#A9B1D6', // LAVENDER / MUTED
  fontWeight: focused ? 'bold' : 'normal',
} as const)

const buttonStyle = (focused: boolean): React.CSSProperties => ({
  color: focused ? '#BB9AF7' : '#A9B1D6', // LAVENDER / MUTED
  fontWeight: focused ? 'bold' : 'normal',
  background: 'transparent',
  border: 'none',
  font: 'inherit',
  padding: '0 0.5rem',
  cursor: 'pointer',
  textAlign: 'left',
  WebkitTapHighlightColor: 'transparent',
} as const)

const errorStyle: React.CSSProperties = { color: '#F7768E', marginTop: '0.75rem' } // RED
const successStyle: React.CSSProperties = { color: '#73DACA', marginTop: '0.75rem' } // GREEN

export function ContactView() {
  const nameRef = useRef<HTMLInputElement>(null)
  const emailRef = useRef<HTMLInputElement>(null)
  const messageRef = useRef<HTMLTextAreaElement>(null)
  const sendRef = useRef<HTMLButtonElement>(null)

  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [message, setMessage] = useState('')
  const [field, setField] = useState(FIELD_NAME)
  const [error, setError] = useState('')
  const [success, setSuccess] = useState(false)
  const [submitting, setSubmitting] = useState(false)

  // Focus the currently active field whenever it changes (mirrors focusContactField).
  useEffect(() => {
    const key = ['name', 'email', 'message', 'send'][field]
    const ref =
      key === 'name'
        ? nameRef
        : key === 'email'
          ? emailRef
          : key === 'message'
            ? messageRef
            : sendRef
    ref?.current?.focus()
  }, [field, nameRef, emailRef, messageRef, sendRef])

  const cycleField = (next: number) => {
    setField(((next % 4) + 4) % 4)
  }

  const handleFieldKeyDown = (e: React.KeyboardEvent, idx: number) => {
    if (e.key === 'Tab' && e.shiftKey) {
      e.preventDefault()
      cycleField(idx - 1)
    } else if (e.key === 'Tab') {
      e.preventDefault()
      cycleField(idx + 1)
    } else if (e.key === 'Enter') {
      e.preventDefault()
      if (idx < FIELD_SEND) {
        cycleField(idx + 1)
      } else {
        handleSend()
      }
    }
  }

  const handleSend = async () => {
    setError('')
    setSuccess(false)

    const trimmedName = name.trim()
    const trimmedEmail = email.trim()
    const trimmedMessage = message.trim()

    if (trimmedName === '') {
      setError('Name is required.')
      return
    }
    if (trimmedEmail === '') {
      setError('Email is required.')
      return
    }
    if (trimmedMessage === '') {
      setError('Message is required.')
      return
    }

    setSubmitting(true)
    try {
      const response = await fetch('/api/v1/contact', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name: trimmedName, email: trimmedEmail, message: trimmedMessage }),
      })

      const data = await response
        .json()
        .catch(() => ({ message: response.statusText }))

      if (!response.ok) {
        setError(data.message || 'Failed to send message.')
        return
      }

      // Mirror the terminal's success message exactly.
      setSuccess(true)
    } catch (err) {
      setError(
        err instanceof Error ? err.message : 'Unable to send the message.',
      )
    } finally {
      setSubmitting(false)
    }
  }

  const indicator = (idx: number) => (field === idx ? '> ' : '  ')

  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100%',
        padding: '1rem',
        boxSizing: 'border-box',
      }}
    >
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          gap: '0.5rem',
          width: '100%',
          maxWidth: 480,
        }}
      >
        {/*
          Name field
        */}
        <div
          onMouseEnter={() => setField(FIELD_NAME)}
          style={{ display: 'flex', flexDirection: 'column', gap: '0.2rem' }}
        >
          <div style={{ display: 'flex', alignItems: 'center', gap: '0.2rem' }}>
            <span style={indicatorStyle(field === FIELD_NAME)}>{indicator(FIELD_NAME)}</span>
            <span style={labelStyle(field === FIELD_NAME)}>Name</span>
          </div>
          <div style={{ paddingLeft: '1rem' }}>
            <input
              ref={nameRef}
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="Your name"
              onKeyDown={(e) => handleFieldKeyDown(e, FIELD_NAME)}
              style={{
                width: '100%',
                background: 'transparent',
                color: '#FFFFFF',
                border: 'none',
                outline: 'none',
                font: 'inherit',
                padding: '0.25rem 0',
              }}
            />
          </div>
        </div>

        {/*
          Email field
        */}
        <div
          onMouseEnter={() => setField(FIELD_EMAIL)}
          style={{ display: 'flex', flexDirection: 'column', gap: '0.2rem' }}
        >
          <div style={{ display: 'flex', alignItems: 'center', gap: '0.2rem' }}>
            <span style={indicatorStyle(field === FIELD_EMAIL)}>{indicator(FIELD_EMAIL)}</span>
            <span style={labelStyle(field === FIELD_EMAIL)}>Email</span>
          </div>
          <div style={{ paddingLeft: '1rem' }}>
            <input
              ref={emailRef}
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="you@example.com"
              onKeyDown={(e) => handleFieldKeyDown(e, FIELD_EMAIL)}
              style={{
                width: '100%',
                background: 'transparent',
                color: '#FFFFFF',
                border: 'none',
                outline: 'none',
                font: 'inherit',
                padding: '0.25rem 0',
              }}
            />
          </div>
        </div>

        {/*
          Message field
        */}
        <div
          onMouseEnter={() => setField(FIELD_MESSAGE)}
          style={{ display: 'flex', flexDirection: 'column', gap: '0.2rem' }}
        >
          <div style={{ display: 'flex', alignItems: 'center', gap: '0.2rem' }}>
            <span style={indicatorStyle(field === FIELD_MESSAGE)}>{indicator(FIELD_MESSAGE)}</span>
            <span style={labelStyle(field === FIELD_MESSAGE)}>Message</span>
          </div>
          <div style={{ paddingLeft: '1rem' }}>
            <textarea
              ref={messageRef}
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              placeholder="Your message..."
              rows={4}
              onKeyDown={(e) => handleFieldKeyDown(e, FIELD_MESSAGE)}
              style={{
                width: '100%',
                background: 'transparent',
                color: '#FFFFFF',
                border: 'none',
                outline: 'none',
                resize: 'vertical',
                font: 'inherit',
                minHeight: '6rem',
              }}
            />
          </div>
        </div>

        <div
          onMouseEnter={() => setField(FIELD_SEND)}
          style={{ display: 'flex', alignItems: 'center', gap: '0.2rem', marginTop: '0.3rem' }}
        >
          <span style={indicatorStyle(field === FIELD_SEND)}>{indicator(FIELD_SEND)}</span>
          <button
            ref={sendRef}
            onKeyDown={(e) => handleFieldKeyDown(e, FIELD_SEND)}
            onClick={handleSend}
            disabled={submitting}
            style={buttonStyle(field === FIELD_SEND)}
          >
            {submitting ? 'Sending...' : '[ Send Message ]'}
          </button>
        </div>

        {error && <div style={errorStyle}>{error}</div>}
        {success && <div style={successStyle}>Thank u for da rose</div>}
      </div>
    </div>
  )
}
