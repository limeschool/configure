import"./index.dcee45e5.js";/* empty css              */import{d as C,e as s,w as g,b1 as D,b2 as V,b9 as h,bH as E,aC as U,aJ as r,aG as F,A as k,aK as l,ba as x}from"./arco.5cfe922d.js";import{C as y,U as q}from"./api.4f3f3ec2.js";const M=C({__name:"form",props:{data:{}},emits:["refresh"],setup(p,{expose:f,emit:c}){const i=s(),t=s(!1),o=s(!1),v=p,u=s({}),w=c;g(()=>v.data,n=>{u.value=n}),f({showAddDrawer:()=>{t.value=!0,o.value=!0},showUpdateDrawer:()=>{t.value=!0,o.value=!1},closeDrawer:()=>{t.value=!1}});const _=async()=>{if(await i.value.validate())return!1;const e=u.value;return o.value?(await y(e),F.success("\u521B\u5EFA\u6210\u529F")):(await q(e),F.success("\u66F4\u65B0\u6210\u529F")),w("refresh"),!0};return(n,e)=>{const m=D,d=x,b=V,A=h,B=E;return k(),U(B,{visible:t.value,"onUpdate:visible":e[3]||(e[3]=a=>t.value=a),title:o.value?"\u65B0\u5EFA":"\u4FEE\u6539",width:"380px",onCancel:e[4]||(e[4]=a=>t.value=!1),onBeforeOk:_},{default:r(()=>[l(A,{ref_key:"formRef",ref:i,model:u.value,"label-align":"left",layout:"horizontal","auto-label-width":""},{default:r(()=>[l(d,{field:"name",label:"\u73AF\u5883\u540D\u79F0",rules:[{required:!0,message:"\u73AF\u5883\u540D\u79F0\u662F\u5FC5\u586B\u9879"}],"validate-trigger":["change","input"]},{default:r(()=>[l(m,{modelValue:u.value.name,"onUpdate:modelValue":e[0]||(e[0]=a=>u.value.name=a),placeholder:"\u8BF7\u8F93\u5165\u73AF\u5883\u540D\u79F0","allow-clear":""},null,8,["modelValue"])]),_:1}),l(d,{field:"keyword",label:"\u73AF\u5883\u6807\u8BC6",rules:[{required:!0,message:"\u73AF\u5883\u6807\u8BC6\u662F\u5FC5\u586B\u9879"}],"validate-trigger":["change","input"]},{default:r(()=>[l(m,{modelValue:u.value.keyword,"onUpdate:modelValue":e[1]||(e[1]=a=>u.value.keyword=a),placeholder:"\u8BF7\u8F93\u5165\u73AF\u5883\u6807\u8BC6","allow-clear":""},null,8,["modelValue"])]),_:1}),l(d,{field:"description",label:"\u73AF\u5883\u63CF\u8FF0",rules:[{required:!0,message:"\u73AF\u5883\u63CF\u8FF0\u662F\u5FC5\u586B\u9879"}],"validate-trigger":["change","input"]},{default:r(()=>[l(b,{modelValue:u.value.description,"onUpdate:modelValue":e[2]||(e[2]=a=>u.value.description=a),placeholder:"\u8BF7\u8F93\u5165\u73AF\u5883\u63CF\u8FF0","allow-clear":""},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1},8,["visible","title"])}}});export{M as _};