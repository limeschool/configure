import{v as r}from"./index.ddbbfbf8.js";function n(e){return r.get("/configure/api/v1/servers",{params:e})}function u(e){return r.post("/configure/api/v1/server",e)}function a(e){return r.put("/configure/api/v1/server",e)}function i(e){return r.delete("/configure/api/v1/server",{params:e})}function s(e){return r.put("/configure/api/v1/server/status",e)}export{u as C,i as D,n as L,s as U,a};