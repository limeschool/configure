import{_ as K}from"./index.ddbbfbf8.js";/* empty css               *//* empty css               */import{d as O,e as r,w as P,b1 as U,aa as X,bh as H,by as Q,b4 as W,ac as Y,bV as Z,c1 as ee,bg as se,b$ as ae,B as n,E as e,aK as s,aJ as c,aO as m,aL as b,aD as te,bc as oe,bi as ne,bj as ce,aG as le,aX as ie,b_ as re,A as o,aM as f,bE as S,aC as V,aN as u}from"./arco.b1c3a1f5.js";import{L as ue}from"./api.9b8fb919.js";import{L as _e}from"./api.09b56d55.js";import"./chart.a760e4f5.js";import"./vue.708d2434.js";const _=i=>(ne("data-v-40766168"),i=i(),ce(),i),de={class:"fields"},pe=_(()=>e("div",{class:"hr"},[e("span",{class:"icon"}),e("span",null,"\u53EF\u7528\u53D8\u91CF")],-1)),ve={class:"header"},me={class:"field-list"},he={class:"item-row"},ye=_(()=>e("span",{class:"label"},"\u5B57\u6BB5",-1)),ge={style:{"margin-right":"5px"}},be={class:"item-row"},fe=_(()=>e("span",{class:"label"},"\u8BF4\u660E",-1)),we={class:"value"},Be={key:0,class:"empty"},ke={class:"item-row"},Ie=_(()=>e("span",{class:"label"},"\u5B57\u6BB5",-1)),Ce={class:"keyword"},De={class:"item-row"},Fe=_(()=>e("span",{class:"label"},"\u8BF4\u660E",-1)),Ee={class:"value"},xe={style:{"margin-right":"5px"}},ze=O({__name:"right",props:{serverId:{type:Number,default:0}},setup(i){const d=i,p=r("business"),v=r(0),w=r([]),h=r([]),l=r({page:1,pageSize:10,keyword:""}),B=a=>`\${${a}}`,L=async()=>{const{data:a}=await _e({...l.value,serverId:d.serverId});w.value=a.list,v.value=a.total},k=async()=>{const{data:a}=await ue({...l.value,serverId:d.serverId});h.value=a.list,v.value=a.total},I=()=>{if(!d.serverId){le.error("\u8BF7\u5148\u9009\u62E9\u670D\u52A1");return}p.value==="business"?k():L()},N=a=>{l.value.page=a,I()},C=()=>{l.value.page=1,l.value.pageSize=10,v.value=0,I()};return P(()=>d.serverId,a=>{a&&k()}),(a,D)=>{const R=U,A=X,G=H,T=Q,F=W,$=oe,E=Y,x=Z,M=ie("svgIcon"),j=ee,q=se,J=ae,z=re("copy");return o(),n("div",de,[pe,e("div",ve,[s(T,null,{default:c(()=>[s(R,{style:{width:"180px"},placeholder:"\u8BF7\u8F93\u5165\u53D8\u91CF\u540D\u79F0","allow-clear":""}),s(G,{type:"primary",onClick:C},{icon:c(()=>[s(A)]),_:1})]),_:1})]),s($,{modelValue:p.value,"onUpdate:modelValue":D[0]||(D[0]=t=>p.value=t),type:"button",style:{width:"100%",didplay:"flex"},onChange:C},{default:c(()=>[s(F,{class:"tag-item",value:"business"},{default:c(()=>[f("\u4E1A\u52A1\u53D8\u91CF")]),_:1}),s(F,{class:"tag-item",value:"resource"},{default:c(()=>[f("\u8D44\u6E90\u53D8\u91CF")]),_:1})]),_:1},8,["modelValue"]),e("div",me,[p.value=="business"?(o(),n(m,{key:0},[(o(!0),n(m,null,b(h.value,(t,y)=>(o(),n("div",{key:y,class:"item"},[e("div",he,[ye,S((o(),V(x,{size:"small",color:"arcoblue"},{default:c(()=>[e("span",ge,u(t.keyword),1),s(E)]),_:2},1024)),[[z,B(t.keyword)]])]),e("div",be,[fe,e("span",we,u(t.description),1)])]))),128)),h.value.length?te("",!0):(o(),n("div",Be,[s(j,null,{image:c(()=>[s(M,{name:"empty-record",size:120})]),default:c(()=>[f(" \u6682\u65E0\u53EF\u7528\u5B57\u6BB5 ")]),_:1})]))],64)):(o(!0),n(m,{key:1},b(w.value,(t,y)=>(o(),n("div",{key:y,class:"item"},[e("div",ke,[Ie,e("span",Ce,u(t.keyword),1)]),e("div",De,[Fe,e("span",Ee,u(t.description),1)]),s(q),(o(!0),n(m,null,b(t.fields.split(","),g=>(o(),n("div",{key:g,class:"item-left"},[S((o(),V(x,{size:"small",color:"arcoblue"},{default:c(()=>[e("span",xe,u(t.keyword+"."+g),1),s(E)]),_:2},1024)),[[z,B(t.keyword+"."+g)]])]))),128))]))),128))]),s(J,{size:"small",total:v.value,current:l.value.page,"page-size":l.value.pageSize,"show-total":"",onChange:N},null,8,["total","current","page-size"])])}}});const $e=K(ze,[["__scopeId","data-v-40766168"]]);export{$e as default};