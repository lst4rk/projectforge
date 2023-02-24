(()=>{function c(e,n){let t;n?t=n.querySelectorAll(e):t=document.querySelectorAll(e);let o=[];return t.forEach(r=>{o.push(r)}),o}function y(e,n){let t=c(e,n);switch(t.length){case 0:return;case 1:return t[0];default:console.warn(`found [${t.length}] elements with selector [${e}], wanted zero or one`)}}function h(e,n){let t=y(e,n);if(!t)throw new Error(`no element found for selector [${e}]`);return t}function T(e,n,t="block"){return typeof e=="string"&&(e=h(e)),e.style.display=n?t:"none",e}function O(){for(let e of c(".menu-container .final"))e.scrollIntoView({block:"nearest"})}var x="mode-light",H="mode-dark";function A(){for(let e of c(".mode-input"))e.onclick=()=>{switch(e.value){case"":document.body.classList.remove(x),document.body.classList.remove(H);break;case"light":document.body.classList.add(x),document.body.classList.remove(H);break;case"dark":document.body.classList.remove(x),document.body.classList.add(H);break}}}function j(e){setTimeout(()=>{e.style.opacity="0",setTimeout(()=>e.remove(),500)},5e3)}function $(e,n,t){let o=document.getElementById("flash-container");o===null&&(o=document.createElement("div"),o.id="flash-container",document.body.insertAdjacentElement("afterbegin",o));let r=document.createElement("div");r.className="flash";let i=document.createElement("input");i.type="radio",i.style.display="none",i.id="hide-flash-"+e,r.appendChild(i);let s=document.createElement("label");s.htmlFor="hide-flash-"+e;let u=document.createElement("span");u.innerHTML="\xD7",s.appendChild(u),r.appendChild(s);let m=document.createElement("div");m.className="content flash-"+n,m.innerText=t,r.appendChild(m),o.appendChild(r),j(r)}function q(){let e=document.getElementById("flash-container");if(e===null)return $;let n=e.querySelectorAll(".flash");if(n.length>0)for(let t of n)j(t);return $}function R(){for(let e of c(".link-confirm"))e.onclick=()=>{let n=e.dataset.message;return n&&n.length===0&&(n="Are you sure?"),confirm(n)}}function w(e,n){let t=(e||"").replace(/-/gu,"/").replace(/[TZ]/gu," ")+" UTC",o=new Date(t),r=(new Date().getTime()-o.getTime())/1e3,i=Math.floor(r/86400),s=o.getFullYear(),u=o.getMonth()+1,m=o.getDate();if(isNaN(i)||i<0||i>=31)return s.toString()+"-"+(u<10?"0"+u.toString():u.toString())+"-"+(m<10?"0"+m.toString():m.toString());let a,d;return i===0?r<5?(d=1,a="just now"):r<60?(d=1,a=Math.floor(r)+" seconds ago"):r<120?(d=10,a="1 minute ago"):r<3600?(d=30,a=Math.floor(r/60)+" minutes ago"):r<7200?(d=60,a="1 hour ago"):(d=60,a=Math.floor(r/3600)+" hours ago"):i===1?(d=600,a="yesterday"):i<7?(d=600,a=i+" days ago"):(d=6e3,a=Math.ceil(i/7)+" weeks ago"),n&&(n.innerText=a,setTimeout(()=>w(e,n),d*1e3)),a}function W(){return c(".reltime").forEach(e=>{w(e.dataset.time||"",e)}),w}function ie(e,n){let t=0;return(...o)=>{t!==0&&window.clearTimeout(t),t=window.setTimeout(()=>{e(null,...o)},n)}}function se(e,n,t,o,r){if(!e)return;let i=e.id+"-list",s=document.createElement("datalist"),u=document.createElement("option");u.value="",u.innerText="Loading...",s.appendChild(u),s.id=i,e.parentElement?.prepend(s),e.setAttribute("autocomplete","off"),e.setAttribute("list",i);let m={},a="";function d(l){let f=n;return f.includes("?")?f+"&t=json&"+t+"="+encodeURIComponent(l):f+"?t=json&"+t+"="+encodeURIComponent(l)}function D(l){let f=m[l];!f||!f.frag||(a=l,s.replaceChildren(f.frag.cloneNode(!0)))}function oe(){let l=e.value;if(l.length===0)return;let f=d(l),b=!l||!a;if(!b){let p=m[a];p&&(b=!p.data.find(k=>l===r(k)))}if(!!b){if(m[l]&&m[l].url===f){D(l);return}fetch(f,{credentials:"include"}).then(p=>p.json()).then(p=>{if(!p)return;let k=Array.isArray(p)?p:[p],N=document.createDocumentFragment(),U=10;for(let v=0;v<k.length&&U>0;v++){let F=r(k[v]),re=o(k[v]);if(F){let L=document.createElement("option");L.value=F,L.innerText=re,N.appendChild(L),U--}}m[l]={url:f,data:k,frag:N,complete:!1},D(l)})}}e.oninput=ie(oe,250),console.log("managing ["+e.id+"] autocomplete")}function _(){return se}function B(){document.addEventListener("keydown",e=>{e.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}function I(e,n,t){return n||(n=18),t==null&&(t="icon"),`<svg class="${t}" style="width: ${n}px; height: ${n+"px"};"><use xlink:href="#svg-${e}"></use></svg>`}function ce(e,n){return e.parentElement!==n.parentElement?null:e===n?0:e.compareDocumentPosition(n)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var M;function S(e){let n=[],t=c(".item .value",e);for(let r of t)n.push(r.innerText);let o=h("input.result",e);o.value=n.join(", ")}function ae(e){let n=e.parentElement?.parentElement;e.remove(),n&&S(n)}function G(e){let n=h(".value",e),t=h(".editor",e);t.value=n.innerText;let o=()=>{if(t.value===""){e.remove();return}n.innerText=t.value,T(n,!0),T(t,!1);let r=e.parentElement?.parentElement;r&&S(r)};t.onblur=o,t.onkeydown=r=>{if(r.code==="Enter")return r.preventDefault(),o(),!1},T(n,!1),T(t,!0),t.focus()}function P(e,n){let t=document.createElement("div");t.className="item",t.draggable=!0,t.ondragstart=s=>{s.dataTransfer?.setDragImage(document.createElement("div"),0,0),t.classList.add("dragging"),M=t},t.ondragover=()=>{let s=ce(t,M);if(!s)return;let u=s===-1?t:t.nextSibling;M.parentElement?.insertBefore(M,u),S(n)},t.ondrop=s=>{s.preventDefault()},t.ondragend=s=>{t.classList.remove("dragging"),s.preventDefault()};let o=document.createElement("div");o.innerText=e,o.className="value",o.onclick=()=>{G(t)},t.appendChild(o);let r=document.createElement("input");r.className="editor",t.appendChild(r);let i=document.createElement("div");return i.innerHTML=I("times",13),i.className="close",i.onclick=()=>{ae(t)},t.appendChild(i),t}function le(e,n){let t=P("",n);e.appendChild(t),G(t)}function J(e){let n=h("input.result",e),t=h(".tags",e),o=n.value.split(",").map(i=>i.trim()).filter(i=>i!=="");T(n,!1),t.innerHTML="";for(let i of o)t.appendChild(P(i,e));y(".add-item",e)?.remove();let r=document.createElement("div");r.className="add-item",r.innerHTML=I("plus",22),r.onclick=()=>{le(t,e)},e.insertBefore(r,h(".clear",e))}function Y(){for(let e of c(".tag-editor"))J(e);return J}var K="--selected";function me(e){let n=e.parentElement?.parentElement?.querySelector("input");if(!n)throw new Error("no associated input found");n.value="\u2205"}function C(e){e.onreset=()=>C(e);let n={},t={};for(let o of e.elements){let r=o;if(r.name.length>0)if(r.name.endsWith(K))t[r.name]=r;else{(r.type!=="radio"||r.checked)&&(n[r.name]=r.value);let i=()=>{let s=t[r.name+K];s&&(s.checked=n[r.name]!==r.value)};r.onchange=i,r.onkeyup=i}}}function Q(){for(let e of c("form.editor"))C(e);return[me,C]}var ue=[];function V(e,n,t){let o=e.querySelectorAll(n);if(o.length===0)throw new Error("empty query selector ["+n+"]");o.forEach(r=>t(r))}function E(e,n,t){V(e,n,o=>{o.style.backgroundColor=t})}function g(e,n,t){V(e,n,o=>{o.style.color=t})}function de(e,n,t){let o=document.querySelector("#mockup-"+e);if(!o){console.error("can't find mockup for mode ["+e+"]");return}switch(n){case"color-foreground":g(o,".mock-main",t);break;case"color-background":E(o,".mock-main",t);break;case"color-foreground-muted":g(o,".mock-main .mock-muted",t);break;case"color-background-muted":E(o,".mock-main .mock-muted",t);break;case"color-link-foreground":g(o,".mock-main .mock-link",t);break;case"color-link-visited-foreground":g(o,".mock-main .mock-link-visited",t);break;case"color-nav-foreground":g(o,".mock-nav",t),g(o,".mock-nav .mock-link",t);break;case"color-nav-background":E(o,".mock-nav",t);break;case"color-menu-foreground":g(o,".mock-menu",t),g(o,".mock-menu .mock-link",t);break;case"color-menu-background":E(o,".mock-menu",t);break;case"color-menu-selected-foreground":g(o,".mock-menu .mock-link-selected",t);break;case"color-menu-selected-background":E(o,".mock-menu .mock-link-selected",t);break;default:console.error("invalid key ["+n+"]")}}function Z(){for(let e of c(".color-var")){let n=e.dataset.var,t=e.dataset.mode;ue.push(n),!(!n||n.length===0)&&(e.oninput=()=>{de(t,n,e.value)})}}var X=!1;function z(e){if(e||(e="/connect"),e.indexOf("ws")===0)return e;let n=document.location,t="ws";return n.protocol==="https:"&&(t="wss"),e.indexOf("/")!==0&&(e="/"+e),t+`://${n.host}${e}`}var ee=class{constructor(n,t,o,r,i){this.debug=n,this.open=t,this.recv=o,this.err=r,this.url=z(i),this.connected=!1,this.pauseSeconds=1,this.pendingMessages=[],this.connect()}connect(){window.onbeforeunload=()=>{X=!0},this.connectTime=Date.now(),this.sock=new WebSocket(z(this.url));let n=this;this.sock.onopen=()=>{n.connected=!0,n.pendingMessages.forEach(n.send),n.pendingMessages=[],n.debug&&console.log("WebSocket connected"),n.open()},this.sock.onmessage=t=>{let o=JSON.parse(t.data);n.debug&&console.debug("[socket]: receive",o),n.recv(o)},this.sock.onerror=t=>()=>{n.err("socket",t.type)},this.sock.onclose=()=>{if(X)return;n.connected=!1;let t=n.connectTime?Date.now()-n.connectTime:0;t>0&&t<2e3?(n.pauseSeconds=n.pauseSeconds*2,n.debug&&console.debug(`socket closed immediately, reconnecting in ${n.pauseSeconds} seconds`),setTimeout(()=>{n.connect()},n.pauseSeconds*1e3)):(console.debug("socket closed after ["+t+"ms]"),setTimeout(()=>{n.connect()},n.pauseSeconds*500))}}disconnect(){}send(n){if(this.debug&&console.debug("out",n),!this.sock)throw new Error("socket not initialized");if(this.connected){let t=JSON.stringify(n,null,2);this.sock.send(t)}else this.pendingMessages.push(n)}};function te(){return ee}function ne(){}function fe(){let[e,n]=Q();window.projectforge={relativeTime:W(),autocomplete:_(),setSiblingToNull:e,initForm:n,flash:q(),tags:Y(),Socket:te()},O(),A(),R(),B(),Z(),ne()}document.addEventListener("DOMContentLoaded",fe);})();
//# sourceMappingURL=client.js.map
