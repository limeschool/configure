import{v as t}from"./index.ddbbfbf8.js";function r(e){return t.get("/configure/api/v1/template/current",{params:{serverId:e}})}function p(e){return t.post("/configure/api/v1/template/switch",e)}function n(e){return t.post("/configure/api/v1/template",e)}function i(e){return t.post("/configure/api/v1/template/preview",e)}function o(e){return t.get("/configure/api/v1/template",{params:{id:e}})}function u(e){return t.get("/configure/api/v1/templates",{params:{...e}})}function m(e){return t.post("/configure/api/v1/template/compare",e)}export{n as C,o as G,u as L,i as P,p as S,r as a,m as b};