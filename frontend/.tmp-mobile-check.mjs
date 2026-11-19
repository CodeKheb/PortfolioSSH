// Temporary diagnostic: emulate a phone in headless Chromium and check scrolling.
const base = 'http://127.0.0.1:9222'
const res = await fetch(`${base}/json/new?http://127.0.0.1:5199/`, { method: 'PUT' })
const target = await res.json()
const ws = new WebSocket(target.webSocketDebuggerUrl)
let id = 0
const pending = new Map()
ws.onmessage = (e) => {
  const msg = JSON.parse(e.data)
  if (msg.id && pending.has(msg.id)) {
    const { resolve, reject } = pending.get(msg.id)
    pending.delete(msg.id)
    msg.error ? reject(new Error(JSON.stringify(msg.error))) : resolve(msg.result)
  }
}
await new Promise((r) => (ws.onopen = r))
const send = (method, params = {}) =>
  new Promise((resolve, reject) => {
    const msgId = ++id
    pending.set(msgId, { resolve, reject })
    ws.send(JSON.stringify({ id: msgId, method, params }))
  })
const evaluate = async (expression) => {
  const r = await send('Runtime.evaluate', { expression, returnByValue: true, awaitPromise: true })
  if (r.exceptionDetails) throw new Error(JSON.stringify(r.exceptionDetails))
  return r.result.value
}
const sleep = (ms) => new Promise((r) => setTimeout(r, ms))

await send('Page.enable')
await send('Runtime.enable')
await send('Emulation.setDeviceMetricsOverride', {
  width: 390, height: 844, deviceScaleFactor: 3, mobile: true,
})
await send('Emulation.setTouchEmulationEnabled', { enabled: true, maxTouchPoints: 5 })
await send('Emulation.setUserAgentOverride', {
  userAgent:
    'Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1',
})
await send('Page.navigate', { url: 'http://127.0.0.1:5199/' })
await sleep(1500)

const measure = `(() => {
  const cs = (el) => el ? getComputedStyle(el) : null
  const pick = (el) => el ? {
    height: cs(el).height, overflowY: cs(el).overflowY, touchAction: cs(el).touchAction,
    clientHeight: el.clientHeight, scrollHeight: el.scrollHeight, scrollTop: el.scrollTop,
  } : null
  const content = document.querySelector('.terminal-content')
  return {
    innerWidth, innerHeight,
    mq640: matchMedia('(max-width: 640px)').matches,
    html: { ...pick(document.documentElement), scrollY: window.scrollY },
    body: pick(document.body),
    root: pick(document.getElementById('root')),
    terminal: pick(document.querySelector('.terminal')),
    content: pick(content),
  }
})()`
const before = await evaluate(measure)
console.log('BEFORE', JSON.stringify(before, null, 2))

const touch = (type, x, y) =>
  send('Input.dispatchTouchEvent', {
    type,
    touchPoints: type === 'touchEnd' ? [] : [{ x, y }],
  })
await touch('touchStart', 195, 600)
for (let y = 600; y >= 200; y -= 40) {
  await touch('touchMove', 195, y)
  await sleep(16)
}
await touch('touchEnd', 195, 200)
await sleep(600)

const after = await evaluate(
  `({ scrollY: window.scrollY, contentScrollTop: (document.querySelector('.terminal-content')||{}).scrollTop, active: document.activeElement && document.activeElement.tagName })`,
)
console.log('AFTER GESTURE', JSON.stringify(after))

// Also try programmatic scroll + screenshot of the About screen to see structure.
console.log('can scroll document?', await evaluate('document.documentElement.scrollHeight > innerHeight'))
ws.close()
