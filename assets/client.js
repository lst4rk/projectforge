(()=>{function k(){for(let e of Array.from(document.querySelectorAll(".menu-container .final")))e.scrollIntoView({block:"nearest"})}function g(){let e=document.getElementById("flash-container");if(e===null)return;let n=e.querySelectorAll(".flash");n.length>0&&setTimeout(()=>{for(let t of n){let o=t;o.style.opacity="0",setTimeout(()=>o.remove(),500)}},3e3)}var y="--selected";function p(){let e={},n={};for(let t of Array.from(document.querySelectorAll(".editor"))){let o=t,c=()=>{e={},n={};for(let s of o.elements){let r=s;if(r.name.length>0)if(r.name.endsWith(y))n[r.name]=r;else{(r.type!=="radio"||r.checked)&&(e[r.name]=r.value);let a=()=>{let l=n[r.name+y];l&&(l.checked=e[r.name]!==r.value)};r.onchange=a,r.onkeyup=a}}};o.onreset=c,c()}}function E(){if(window.Sortable)for(let e of Array.from(document.getElementsByClassName("sortable")))h(e)}function b(e){for(;e.parentElement&&!e.classList.contains("drag-container");)e=e.parentElement;e.classList.remove("readonly"),h(e)}function h(e){let n=window.Sortable;if(n){let t=e.querySelector(".l");t||(t=e);let c={group:{name:"nested"},handle:".handle",onAdd:r=>{let a=r.item;new n(a.querySelector(".container"),c),a.querySelector(".remove").onclick=function(){L(e,a)},u(e)},onUpdate:()=>u(e),animation:150,fallbackOnBody:!0,swapThreshold:.65};for(let r of Array.from(t.getElementsByClassName("container")))new n(r,c);for(let r of Array.from(t.getElementsByClassName("remove")))r.onclick=function(){L(e,r.parentElement?.parentElement)};let s=e.querySelector(".r");if(s){let r={group:{name:"nested",pull:"clone",put:!1},handle:".handle",animation:150,fallbackOnBody:!0,swapThreshold:.65,sort:!1};for(let a of Array.from(s.getElementsByClassName("container")))new n(a,r)}u(e)}}function L(e,n){n.remove(),u(e)}function u(e){let n=document.querySelector(".drag-state");if(!n)return;let t=document.querySelector(".drag-state-original"),o=e.querySelector(".tracked"),[c,s]=I(o),r=JSON.stringify(c);if(t){t.value.length===0&&(t.value=r);let a=document.querySelector(".drag-actions");a&&(t.value===r?a.classList.add("no-changes"):a.classList.remove("no-changes"));let l=document.querySelector(".drag-tracked-size");l&&(s===1?l.innerText=s.toString(10)+(l.dataset.sing?" "+l.dataset.sing:""):l.innerText=s.toString(10)+(l.dataset.plur?" "+l.dataset.plur:""))}n.value=r}function I(e){let n=0,t=[];for(let o of Array.from(e.children))if(o.classList.contains("item")){let[c,s]=H(o);c&&t.push(c),n+=s}return[t,n]}function H(e){let n=1,t={k:e.dataset.key,p:e.dataset.originalPath};for(let o of Array.from(e.children))if(o.classList.contains("container")){let[c,s]=I(o);c.length>0&&(t.c=c),n+=s}return[t,n]}function T(){for(let e of Array.from(document.getElementsByClassName("link-confirm"))){let n=e;n.onclick=function(){let t=n.dataset.message;return t&&t.length===0&&(t="Are you sure?"),confirm(t)}}}var q=[];function S(){let e=document.querySelectorAll(".color-var");if(e.length>0)for(let n of Array.from(e)){let t=n,o=t.dataset.var,c=t.dataset.mode;q.push(o),!(!o||o.length===0)&&(t.oninput=function(){w(c,o,t.value)})}}function w(e,n,t){let o=document.querySelector("#mockup-"+e);if(!o){console.error("can't find mockup for mode ["+e+"]");return}switch(n){case"color-foreground":i(o,".mock-main",t);break;case"color-background":m(o,".mock-main",t);break;case"color-foreground-muted":i(o,".mock-main .mock-muted",t);break;case"color-background-muted":m(o,".mock-main .mock-muted",t);break;case"color-link-foreground":i(o,".mock-main .mock-link",t);break;case"color-link-visited-foreground":i(o,".mock-main .mock-link-visited",t);break;case"color-nav-foreground":i(o,".mock-nav",t),i(o,".mock-nav .mock-link",t);break;case"color-nav-background":m(o,".mock-nav",t);break;case"color-menu-foreground":i(o,".mock-menu",t),i(o,".mock-menu .mock-link",t);break;case"color-menu-background":m(o,".mock-menu",t);break;case"color-menu-selected-foreground":i(o,".mock-menu .mock-link-selected",t);break;case"color-menu-selected-background":m(o,".mock-menu .mock-link-selected",t);break;default:console.error("invalid key ["+n+"]")}}function v(e,n,t){let o=e.querySelectorAll(n);if(o.length==0)throw"empty query selector ["+n+"]";o.forEach(c=>{t(c)})}function m(e,n,t){v(e,n,o=>o.style.backgroundColor=t)}function i(e,n,t){v(e,n,o=>o.style.color=t)}function A(e){let n=e.parentElement.parentElement.querySelector("input");if(!n)throw"no associated input found";n.value="\u2205"}var f="mode-light",d="mode-dark";function M(){for(let e of Array.from(document.getElementsByClassName("mode-input"))){let n=e;n.onclick=function(){switch(n.value){case"":document.body.classList.remove(f),document.body.classList.remove(d);break;case"light":document.body.classList.add(f),document.body.classList.remove(d);break;case"dark":document.body.classList.remove(f),document.body.classList.add(d);break}}}}function x(){window.projectforge={sortableEdit:b,setSiblingToNull:A},k(),M(),g(),T(),p(),E(),S()}document.addEventListener("DOMContentLoaded",x);})();
//# sourceMappingURL=client.js.map
