const base='http://127.0.0.1:9222'
const t=await (await fetch(`${base}/json/list`)).json()
const page=t.find(x=>x.type==='page')
const ws=new WebSocket(page.webSocketDebuggerUrl)
let id=0;const p=new Map()
ws.onmessage=e=>{const m=JSON.parse(e.data);if(m.id&&p.has(m.id)){const{resolve}=p.get(m.id);p.delete(m.id);resolve(m.result)}}
await new Promise(r=>ws.onopen=r)
const send=(method,params={})=>new Promise(res=>{const i=++id;p.set(i,{resolve:res});ws.send(JSON.stringify({id:i,method,params}))})
const ev=async(x)=>(await send('Runtime.evaluate',{expression:x,returnByValue:true,awaitPromise:true})).result.value
console.log('url',await ev('location.href'))
console.log('title',await ev('document.title'))
console.log('root?',await ev('!!document.getElementById("root")'))
console.log(await ev('document.body.innerHTML.slice(0,500)'))
ws.close()
