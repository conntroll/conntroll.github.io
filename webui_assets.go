// Code generated by assets compiler. DO NOT EDIT.
	
package assets

var (
	Assets = map[string]string{
"/dist/bundle.js": "!function(e){function t(o){if(n[o])return n[o].exports;var r=n[o]={i:o,l:!1,exports:{}};return e[o].call(r.exports,r,r.exports,t),r.l=!0,r.exports}var n={};t.m=e,t.c=n,t.d=function(e,n,o){t.o(e,n)||Object.defineProperty(e,n,{enumerable:!0,get:o})},t.r=function(e){\"undefined\"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:\"Module\"}),Object.defineProperty(e,\"__esModule\",{value:!0})},t.t=function(e,n){if(1&n&&(e=t(e)),8&n)return e;if(4&n&&\"object\"==typeof e&&e&&e.__esModule)return e;var o=Object.create(null);if(t.r(o),Object.defineProperty(o,\"default\",{enumerable:!0,value:e}),2&n&&\"string\"!=typeof e)for(var r in e)t.d(o,r,function(t){return e[t]}.bind(null,r));return o},t.n=function(e){var n=e&&e.__esModule?function(){return e.default}:function(){return e};return t.d(n,\"a\",n),n},t.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},t.p=\"\",t(t.s=0)}([function(){\"use strict\";function e(){let e,t=document.getElementById(\"api\").value;try{e=new URL(t)}catch(e){r(function(e){let t=document.createElement(\"div\"),n=new Date;return t.innerHTML=n.toISOString()+\" \"+e,t}(e))}return e}function t(e){let t={id:e.id,name:e.name,ip:e.ip,os:e.os,pwd:e.pwd,arch:e.arch,auth:e.auth,tags:e.tags,username:e.username,hostname:e.hostname,connected:e.connected};return e.distro&&(t.distro=e.distro),t}function n(t){let n=\"/api/agent/\",o=e(),r=o.protocol+\"//\"+o.host+n+t.id;\"\"!=o.username&&\"\"!=o.password&&(r=o.protocol+\"//\"+o.username+\":\"+o.password+\"@\"+o.host+n+t.id);let i=document.createElement(\"div\"),a=document.createElement(\"div\"),d=document.createElement(\"div\"),l=document.createElement(\"div\"),c=document.createElement(\"div\"),u=document.createElement(\"a\"),s=document.createElement(\"a\"),p=document.createElement(\"a\"),m=document.createElement(\"span\"),h=document.createElement(\"span\"),f=document.createElement(\"span\"),g=document.createElement(\"span\"),b=document.createElement(\"div\"),w=document.createElement(\"a\"),v=document.createElement(\"a\"),E=document.createElement(\"a\");u.innerHTML=t.username+\"@\"+t.hostname+\" \",s.innerHTML=\" \"+t.pwd,s.setAttribute(\"target\",\"_blank\"),s.setAttribute(\"href\",r+\"/rootfs\"+t.pwd),p.innerHTML=\" ~ \",p.setAttribute(\"target\",\"_blank\"),p.setAttribute(\"href\",r+\"/rootfs/home/\"+t.username),m.appendChild(u),h.appendChild(s);for(let e of t.tags){let t=document.createElement(\"a\");t.innerHTML=e,t.setAttribute(\"class\",\"tag\"),f.appendChild(t)}if(g.appendChild(p),E.innerHTML=t.ip,E.setAttribute(\"target\",\"_blank\"),E.setAttribute(\"href\",\"https://www.wolframalpha.com/input/?i=\"+t.ip+\" ip geolocation\"),v.innerHTML=\" | \",w.innerHTML=\"connect\",w.setAttribute(\"target\",\"_blank\"),w.setAttribute(\"href\",r+\"/\"),l.setAttribute(\"class\",\"agent-topbar\"),l.appendChild(m),l.appendChild(f),c.setAttribute(\"class\",\"agent-midbar\"),c.appendChild(h),c.appendChild(g),b.setAttribute(\"class\",\"agent-subbar\"),b.innerHTML=\"connected \"+function(e){var t=Math.round;let n=t((Date.now()/1e3-e)/60);return 60>n?n.toString()+\" min\":60<=n&&1440>n?t(n/60).toString()+\" hour\":t(n/60/24).toString()+\" day\"}(t.connected)+\" ago | \"+t.name+\" | \",b.appendChild(E),b.appendChild(v),b.appendChild(w),a.setAttribute(\"class\",\"agent-left\"),a.className+=t.auth?\" auth\":\" noauth\",a.innerHTML=\"*\",t.distro){let e=t.distro,n=e.slice(0,1).toUpperCase()+e.slice(1,2);a.innerHTML=n}return d.setAttribute(\"class\",\"agent-right\"),d.appendChild(l),d.appendChild(c),d.appendChild(b),i.setAttribute(\"class\",\"agent\"),i.setAttribute(\"id\",t.id),i.appendChild(a),i.appendChild(d),i}function o(e){let t=document.getElementById(\"agents\"),o=new Set,r=new Set;for(let e of t.childNodes)o.add(e.id);for(let t of e)r.add(t.id);for(let r of e)if(!o.has(r.id)){let e=n(r);t.insertBefore(e,t.firstChild),o.add(r.id)}for(let e of t.childNodes)r.has(e.id)||t.removeChild(e)}function r(e){let t=document.getElementById(\"stats\");for(;3<=t.childNodes.length;)t.removeChild(t.lastChild);t.prepend(e)}function i(){let e=\"https:\"==window.location.protocol||s?\"wss://\":\"ws://\",n=s?a:window.location.host;var d=new WebSocket(e+n+\"/api/agents/watch\");d.binaryType=\"blob\",d.onopen=()=>{r(function(){let e=document.createElement(\"div\"),t=new Date;return e.innerHTML=t.toISOString()+\" Start Watching Agents\",e}())},d.onclose=e=>{r(function(){let e=document.createElement(\"div\"),t=new Date;return e.innerHTML=t.toISOString()+\" WebSocket Connection Closed\",e}()),setTimeout((function(){i()}),5)},d.addEventListener(\"message\",(function(e){let n=[],i=JSON.parse(e.data);for(let e of i)n.push(t(e));o(n),r(function(e){let t=document.createElement(\"div\"),n=new Date;return t.innerHTML=n.toISOString()+\" \"+e.toString()+\" Agents Connected\",t}(n.length))}))}this&&this.__awaiter;var a=\"hub.k0s.io\",d=\"https://\"+a,l=window.chrome&&window.chrome.runtime&&window.chrome.runtime.id,c=\"conntroll.github.io\"==window.location.host,u=window.location.host.endsWith(\"k0s.io\"),s=l||c||u;window.onload=function(){(function(){let e=document.getElementById(\"api\");s?e.value=d:e.value=window.location.origin})(),i()}}]);\n//# sourceMappingURL=bundle.js.map",
"/index.html": "<!DOCTYPE html>\n<html>\n <head>\n  <meta http-equiv=\"content-type\" content=\"text/html; charset=utf-8\" />\n  <link href=\"style.css\" rel=\"stylesheet\" type=\"text/css\" />\n  <script src=\"dist/bundle.js\"></script>\n </head>\n <body>\n  <div id=\"main\">\n    <div id=\"controller\">\n      <input id=\"api\" value=\"https://libredot.com\"></input>\n      <button id=\"list\">List</button>\n      <button id=\"watch\">Watch</button>\n    </div>\n    <div id=\"view\">\n\n      <div class=\"doc\" id=\"topbar\">\n          <a id=\"a-story\" href='#'>Story</a> \n        | <a id=\"a-download\" href='#'>Download</a>\n        | <a id=\"a-examples\" href='#'>Examples</a>\n        | <a id=\"a-meaning-of-colors\" href='#'\">Meaning of colors</a> \n        | <a id=\"a-stats\" href='#')>Stats</a>\n        | <a id=\"a-about\" href='#')>About</a>\n        <!-- | <a id=\"a-how-it-works\" href='#'>How it works</a> -->\n      </div>\n\n      <div class=\"doc\" id=\"story\" style=\"display: block;\">\n        <h1>ConnTroll</h1>\n\n        <hr>\n\n        <h2>What is ConnTroll?</h2>\n        <p>ConnTroll was a tool initially designed to control my ThinkPads at home from work. </p>\n        <p>It helped a lot until lately the company I worked for went out of business during the trade war. Consequently I'm out of work, which makes ConnTroll less useful to me than to other comrades. That's the background why I decided to polish and publish it.</p>\n\n        <hr>\n\n        <h2>What does it do?</h2>\n        <p>Common use cases include:</p>\n\t<ul>\n          <li>control home computer from work</li>\n          <li>access terminal and files behind NAT</li>\n          <li>stream container logs to browser in real time</li>\n\t</ul>\n        <p>Please note that the development of ConnTroll is still in an early stage, some of the features may be incomplete or totally disfunctional. Misconfiguration of the tool may lead to your personal data being leaked or your system being compromised. Use it at your own risk.</p>\n\n        <hr>\n\n        <h2>How do I use it?</h2>\n        <p>See the #Examples section</p>\n      </div>\n\n      <div class=\"doc\" id=\"download\">\n        <h2>Download</h2>\n        <!--<p>find below a container named download@conntroll</p>-->\n        <p>download latest binary release (updated 2020-01-01):</p>\n\t<ul>\n          <li><a href=\"https://conntroll.github.io/releases/latest/linux/amd64/agent\">linux/amd64</a></li>\n          <li><a href=\"https://conntroll.github.io/releases/latest/linux/386/agent\">linux/386</a></li>\n          <li><a href=\"https://conntroll.github.io/releases/latest/darwin/amd64/agent\">darwin/amd64</a></li>\n\t</ul>\n        <p>alternatively you can pull the latest docker image via:</p>\n        <pre>$ docker run conntroll/agent:latest</pre>\n      </div>\n\n      <div class=\"doc\" id=\"examples\">\n        <h2>Examples</h2>\n        <h3>Run on your host</h3>\n        <p>Although I used to do so, running conntroll directly on your host is strongly discoraged, unless you know what you are doing</p>\n        <h3>Run inside a container</h3>\n\t<pre>\n$ docker run conntroll/agent \n$ docker run conntroll/agent:arch\n$ docker run conntroll/agent:ubuntu\n$ docker run conntroll/agent:alpine\n$ docker run conntroll/agent:gentoo\n$ docker run conntroll/agent:fedora\n$ docker run conntroll/agent:debian\n\n$ docker run -v /dir/to/share:/public -w /public conntroll/agent \n</pre>\n      </div>\n\n      <div class=\"doc\" id=\"how-it-works\">\n        <h2>How it works</h2>\n        <p>There are three parties involved: hub, agent, and client</p>\n      </div>\n\n      <div class=\"doc\" id=\"meaning-of-colors\">\n        <h2>Meaning of colors</h2>\n\t<fieldset>\n\t  <legend>Legend</legend>\n\t  <div class=\"foobar\">\n\t    <div class=\"foo secret\"></div>\n\t    <label class=\"bar\" for=\"mothman\"> #5EC2E7 Basic Auth</label>\n\t  </div>\n\t  <div class=\"foobar\">\n\t    <div class=\"foo badass\"></div>\n\t    <label class=\"bar\" for=\"sasquatch\"> #BADA55 No Auth</label>\n\t  </div>\n\t</fieldset>\n      </div>\n\n      <div class=\"doc\" id=\"about\">\n        <h2>About</h2>\n        <p>Email: navigaid@gmail.com</p>\n        <p>Wechat: navigaid</p>\n        <p>GitHub: https://github.com/btwiuse/conntroll</p>\n      </div>\n\n      <div class=\"doc\" id=\"stats\"></div>\n\n      <div id=\"agents\"></div>\n\n    </div>\n  </div>\n </body>\n  <script src=\"script.js\"></script>\n</html>\n",
"/script.js": "// function toggleElementVisibilityById(id){\nfunction toggle(id){\n  elem = document.getElementById(id);\n  disp = elem.style.display;\n  if (disp != \"block\") {\n    elem.style.display = \"block\";\n    setCookie(\"prevtab\", id, 7)\n  } else {\n    elem.style.display = \"\";\n  }\n  [\"story\", \"download\", \"examples\", \"meaning-of-colors\", \"stats\", \"about\"].forEach((e)=>{ // \"how-it-works\", \n    if (e != id) {\n      document.getElementById(e).style.display=\"\";\n    }\n  });\n}\n// https://stackoverflow.com/questions/14573223/set-cookie-and-get-cookie-with-javascript\nfunction setCookie(name,value,days) {\n    var expires = \"\";\n    if (days) {\n        var date = new Date();\n        date.setTime(date.getTime() + (days*24*60*60*1000));\n        expires = \"; expires=\" + date.toUTCString();\n    }\n    document.cookie = name + \"=\" + (value || \"\")  + expires + \"; path=/\";\n}\nfunction getCookie(name) {\n    var nameEQ = name + \"=\";\n    var ca = document.cookie.split(';');\n    for(var i=0;i < ca.length;i++) {\n        var c = ca[i];\n        while (c.charAt(0)==' ') c = c.substring(1,c.length);\n        if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);\n    }\n    return null;\n}\nfunction eraseCookie(name) {   \n    document.cookie = name+'=; Max-Age=-99999999;';  \n}\nvar tab = getCookie(\"prevtab\");\nif (tab) {\n  toggle(tab)\n}\n[\"story\", \"download\", \"examples\", \"meaning-of-colors\", \"stats\", \"about\"].forEach((e)=>{ // \"how-it-works\",\n  document.getElementById(\"a-\"+e).onclick=(()=>toggle(e));\n});\n",
"/style.css": "    body {\n      margin: 0px;\n      background: aquamarine;\n    }\n    hr {\n      border-style: dashed;\n    }\n    pre {\n      overflow-wrap: break-word;\n      overflow: auto;\n    }\n    div {\n      overflow-wrap: break-word;\n      text-align: left;\n    }\n    a {\n      text-decoration: none;\n      // color: fuchsia;\n    }\na:link {\n  color: purple;\n}\n\na:hover {\n  color: fuchsia;\n}\n\na:active {\n  color: green;\n}\n\na.tag {\n\tbackground-color: #fffcd7;\n\tborder: 1px solid #d5d458;\n\tborder-radius: 5px;\n\tcolor: #555;\n\tfont-size: 10pt;\n\tmargin-left: 0.25em;\n\tpadding: 0px 0.4em 1px 0.4em;\n\ttext-decoration: none;\n\tvertical-align: text-top;\n}\n\n    #main {\n      min-width: 600px;\n    }\n    #topbar {\n      display: block;\n    }\n    .doc {\n      display: none;\n      color: grey;\n      text-align: left;\n      font-size: 11.5pt;\n      font-weight: bold;\n      font-family: \"Dejavu Sans\", \"helvetica neue\", aria, sans-serif;\n      background: white;\n      border-width: thin;\n      // border-color: black;\n      // border-style: solid;\n      border-radius: 1em;\n      margin-left: 1em;\n      margin-right: 1em;\n      padding-top: 1em;\n      padding-bottom: 1em;\n      padding-left: 3em;\n      padding-right: 3em;\n      margin-top: 1em;\n      margin-bottom: 1em;\n      // box-shadow: 3px 3px 13px 0px black;\n    }\n    #controller{\n      display: none;\n    }\n    .auth {\n      background: #5ec2e7;\n    }\n    .noauth {\n      background: #bada55;\n    }\n    .agent {\n      background: white;\n      border-width: thin;\n      // border-color: black;\n      // border-style: solid;\n      position: relative;\n      border-radius: 1em;\n      margin-left: 1em;\n      margin-right: 1em;\n      margin-top: 1em;\n      margin-bottom: 1em;\n      padding: 1em;\n      // box-shadow: 3px 3px 13px 0px black;\n    }\n    .agent-left{\n      float: left;\n      position: absolute;\n      border-color: gray;\n      border-right-style: dashed;\n      font-family: \"Dejavu Sans\", \"helvetica neue\", aria, sans-serif;\n      border-width: thin;\n      top: 0px;\n      bottom: 0px;\n      width: 2em;\n      left: 0px;\n      border-bottom-left-radius: 1em;\n      border-top-left-radius: 1em;\n      // background-color: peachpuff;\n      // background: #bada55;\n      text-align: center;\n      padding-top: 1.5em;\n      padding-bottom: 1.5em;\n    }\n    .agent-right{\n      margin-left: 2em;\n    }\n    .agent-topbar{\n      font-family: \"Dejavu Sans Mono\", \"helvetica neue\", aria, sans-serif;\n      font-size: 11.5pt;\n      font-weight: bold;\n      line-height: 1.45em;\n      margin-bottom: 0.25em;\n    }\n    .agent-midbar{\n      font-family: \"Dejavu Sans Mono\", \"helvetica neue\", aria, sans-serif;\n      font-size: 11.5pt;\n      font-weight: bold;\n      line-height: 1.45em;\n      margin-bottom: 0.55em;\n    }\n    .agent-subbar{\n      font-family: \"Dejavu Sans\", \"helvetica neue\", aria, sans-serif;\n      font-size: 10.5pt;\n      line-height: 0.85em;\n      color: #888;\n    }\n    .legend {\n      width: 2em;\n      height: 2em;\n      background: black;\n    }\n\n.foo {\n  float: left;\n  width: 20px;\n  height: 20px;\n  margin: 5px;\n  border: 1px solid rgba(0, 0, 0, .2);\n}\n\n.bar {\n  float: left;\n  // width: 8em;\n  height: 20px;\n  margin: 5px;\n  // border: 1px solid rgba(0, 0, 0, .2);\n}\n\n.badass {\n  background: #bada55;\n}\n\n.secret {\n  background: #5ec2e7;\n}\n\n.foobar {\n  position: relative;\n  min-width: 15em;\n  max-width: 18em;\n  width: 40%;\n  float: left;\n  margin-left: 1em;\n}\n\nfieldset {\n// border-style: solid;\nborder-radius: 1em;\nborder-color: pink;\nmargin-bottom: 1em;\n}\n",
}
)