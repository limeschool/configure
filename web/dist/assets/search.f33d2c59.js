var h=Object.defineProperty;var l=(a,s,e)=>s in a?h(a,s,{enumerable:!0,configurable:!0,writable:!0,value:e}):a[s]=e;var r=(a,s,e)=>(l(a,typeof s!="symbol"?s+"":s,e),e);class n{constructor(s,e,t){r(this,"store",{});r(this,"result",[]);r(this,"params",{page:1,pageSize:10});r(this,"isSelectdHandler");r(this,"queryHandler");r(this,"search",async()=>this.queryHandler?await this.queryHandler(this.params):[]);r(this,"Search",async s=>{this.store={};for(let t=0;t<this.result.length;t+=1){const i=this.result[t];this.store[i.value]=!0,(this.isSelectdHandler==null||!this.isSelectdHandler(i.value))&&(this.result.splice(t,1),t-=1,this.store[i.value]=!1)}this.params.page=1,this.params.pageSize=10,this.params.query=s,(await this.search()).forEach(t=>{this.store[t.value]||this.result.push(t)})});r(this,"NextSearch",async()=>{if(!this.queryHandler||this.result.length<this.params.page*this.params.pageSize)return;this.params.page+=1,(await this.search()).forEach(e=>{this.store[e.value]||this.result.push(e)})});r(this,"IsExist",async s=>this.store[s]);this.result=s,this.queryHandler=e,this.isSelectdHandler=t}}export{n as S};