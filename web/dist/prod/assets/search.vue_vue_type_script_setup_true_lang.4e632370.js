import{f as B}from"./index.9de08a74.js";/* empty css               *//* empty css               *//* empty css               */import{d as D,e as f,o as F,aP as V,bm as S,bo as I,aa as E,aX as A,aW as z,be as N,aC as R,aI as a,bt as U,A as $,aJ as e,aL as v,bp as q,bu as J,bv as L}from"./arco.79b3a20c.js";import{p as M}from"./server.cc9bc429.js";const K=D({__name:"search",emits:["search","select"],setup(P,{emit:p}){const t=f({}),c=f([]),h=()=>{p("search",t.value)},b=()=>{t.value.keyword=void 0},d=async r=>{const{data:n}=await M({page:1,page_size:10,keyword:r}),{list:u}=n;if(!u)return;const l=[];u.forEach(o=>{t.value.server_id!==o.id&&(o.name=`${o.name}(${o.keyword})`,l.push(o))});const s=[];c.value.forEach(o=>{o.id===t.value.server_id&&s.push(o)}),c.value=l.concat(s)},w=r=>{p("select",r)};return F(()=>{d()}),(r,n)=>{const u=V,l=q,s=J,o=S,i=U,C=I,y=E,m=A,k=B,g=z,x=L("permission");return N(($(),R(i,null,{default:a(()=>[e(s,{flex:1},{default:a(()=>[e(C,{model:t.value,"label-col-props":{span:6},"wrapper-col-props":{span:18},"label-align":"left","auto-label-width":""},{default:a(()=>[e(i,{gutter:16},{default:a(()=>[e(s,{span:8},{default:a(()=>[e(l,{field:"keyword",label:"\u670D\u52A1\u6807\u8BC6"},{default:a(()=>[e(u,{modelValue:t.value.server_id,"onUpdate:modelValue":n[0]||(n[0]=_=>t.value.server_id=_),placeholder:"\u8BF7\u9009\u62E9\u670D\u52A1",scrollbar:!0,options:c.value,"field-names":{value:"id",label:"name"},onSearch:d,onChange:w},null,8,["modelValue","options"])]),_:1})]),_:1}),e(s,{span:8},{default:a(()=>[e(l,{field:"keyword",label:"\u53D8\u91CF\u6807\u8BC6"},{default:a(()=>[e(o,{modelValue:t.value.keyword,"onUpdate:modelValue":n[1]||(n[1]=_=>t.value.keyword=_),"allow-clear":"",placeholder:"\u8BF7\u8F93\u5165\u53D8\u91CF\u6807\u8BC6"},null,8,["modelValue"])]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(s,{flex:"86px",style:{"text-align":"right"}},{default:a(()=>[e(g,{size:18},{default:a(()=>[e(m,{type:"primary",onClick:h},{icon:a(()=>[e(y)]),default:a(()=>[v(" \u641C\u7D22 ")]),_:1}),e(m,{onClick:b},{icon:a(()=>[e(k)]),default:a(()=>[v(" \u91CD\u7F6E ")]),_:1})]),_:1})]),_:1})]),_:1})),[[x,"configure:business:query"]])}}});export{K as _};