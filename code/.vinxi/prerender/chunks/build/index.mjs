import { createComponent, ssr, ssrHydrationKey, escape, ssrAttribute } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/solid-js/web/dist/server.js';
import { createContext, createSignal, createEffect, useContext } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/solid-js/dist/server.js';
import { makePersisted } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/@solid-primitives/storage/dist/index.js';
import { createStore } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/solid-js/store/dist/server.js';
import { v as ve$1, a as v, E, N as N$1, p as pe$1, w as we$1, b as be$1, g as ge$1, U as U$1, c as p, f as fe$1, y as ye$1, d as ce$1 } from './game-board-b1526277.mjs';
import { debounce } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/@solid-primitives/scheduled/dist/index.js';
import { f } from './index-d741c3e0.mjs';
import 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/seedrandom/index.js';

const M = createContext();
function F(e) {
  let [l, r] = makePersisted(createStore({ dialog_status: true }), { name: "mural_info-dialog" });
  const a = [l, { close() {
    r("dialog_status", false), document.body.style.overflowY = "auto";
  }, open() {
    r("dialog_status", true), document.body.style.position = "relative", document.body.style.overflowY = "hidden";
  } }];
  return createComponent(M.Provider, { value: a, get children() {
    return e.children;
  } });
}
function D() {
  return useContext(M);
}
var G = ["<button", ' id="info-button" class=" w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700 ">Info</button>'];
function O() {
  return D(), ssr(G, ssrHydrationKey());
}
var q = ["<div", ' class="', '">Submit</div>'];
function R() {
  const [e, l] = v();
  return ssr(q, ssrHydrationKey(), `${e.selected_option ? "" : "bg-river-bed-400"} ${e.selected_option ? "border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700" : ""}  w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center `);
}
var U = ["<button", ' class="w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center hover:cursor-pointer bg-contessa-600">Give Up</button>'];
function W() {
  return v(), ssr(U, ssrHydrationKey());
}
var V = ["<button", ' class="', '"><!--$-->', "<!--/--> (<!--$-->", "<!--/-->)</button>"];
function K(e) {
  v();
  const l = e.disabled, r = e.movie;
  return ssr(V, ssrHydrationKey(), `${l ? "" : "hover:cursor-pointer"} text-left text-lg w-full rounded-lg py-2 px-4 bg-transparent text-river-bed-700 border-2 border-river-bed-700 `, escape(r.title), escape(new Date(r.release_date).getFullYear()));
}
var N = ["<div", ' class="', '"><!--$-->', "<!--/--> (<!--$-->", "<!--/-->)</div>"];
function Z(e) {
  const l = e.disabled, r = e.movie;
  return ssr(N, ssrHydrationKey(), `${l ? "" : "hover:cursor-pointer"} text-lg w-full rounded-lg bg-desert-sand-300 text-river-bed-700 border-4 border-river-bed-800 py-2 px-4`, escape(r.title), escape(new Date(r.release_date).getFullYear()));
}
var J = ["<ul", ' class="w-full flex flex-col space-y-1">', "</ul>"];
function Q() {
  const [e, l] = v();
  return ssr(J, ssrHydrationKey(), escape(e.easy_mode_options.map((r) => {
    var _a, _b, _c, _d;
    return ((_a = e.selected_option) == null ? void 0 : _a.id) == r.id ? r.id == ((_b = e.correct_option) == null ? void 0 : _b.id) && e.status != p.init && e.status != p.started ? createComponent(fe$1, { disabled: true, movie: r }) : r.id != ((_c = e.correct_option) == null ? void 0 : _c.id) && e.status != p.init && e.status != p.started ? createComponent(ye$1, { disabled: true, movie: r }) : createComponent(Z, { get disabled() {
      return !(e.status != p.init && e.status != p.started);
    }, movie: r }) : r.id == ((_d = e.correct_option) == null ? void 0 : _d.id) && e.status != p.init && e.status != p.started ? createComponent(fe$1, { disabled: true, movie: r }) : createComponent(K, { get disabled() {
      return e.status != p.init && e.status != p.started;
    }, movie: r });
  })));
}
var X = ["<div", ' class="flex flex-col space-y-2"><!--$-->', "<!--/--><!--$-->", "<!--/--><!--$-->", "<!--/--><div></div></div>"], ee = ["<div", ' class="relative"><div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none"><svg class="w-4 h-4 text-river-bed-700" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"></path></svg></div><input id="search-query" name="search-query" placeholder="Enter a movie title..."', ' class="block w-full px-4 py-2 ps-10 text-sm text-river-bed-700 bg-desert-sand-100 border-2 border-river-bed-700 placeholder:text-river-bed-700 rounded-md focus:ring-river-bed-700 focus:border-river-bed-700" required></div>'], te = ["<div", ' class="z-10 bg-desert-sand-100 border-2 border-river-bed-700 divide-y divide-desert-sand-100 rounded-lg shadow w-full overflow-scroll no-scrollbar no-scrollbar::-webkit-scrollbar max-h-44" id="answer-options"><ul class="py-2 text-sm text-river-bed-700">', "</ul></div>"], re = ["<li", ' class="block px-4 py-2 hover:bg-desert-sand-200 hover:cursor-pointer"><!--$-->', "<!--/--> (<!--$-->", "<!--/-->)</li>"], se = ["<div", ">", "</div>"], de = ["<div", ' class="flex flex-col space-y-2"><!--$-->', "<!--/--><!--$-->", "<!--/--></div>"];
function ie() {
  var _a, _b, _c;
  const [e, l] = v(), [r, a] = createSignal([]);
  return debounce((c) => {
    if (c != "") {
      const f = ce$1(c).slice(0, 10);
      a([...f]);
    } else
      a([]);
  }, 300), ssr(X, ssrHydrationKey(), e.status == p.started ? ssr(ee, ssrHydrationKey(), ssrAttribute("value", e.selected_option ? `${escape((_a = e.selected_option) == null ? void 0 : _a.title, true)} (${escape(new Date((_b = e.selected_option) == null ? void 0 : _b.release_date).getFullYear(), true)}) (${escape((_c = e.selected_option) == null ? void 0 : _c.id, true)})` : "", false)) : escape([]), r().length >= 1 ? ssr(te, ssrHydrationKey(), escape(r().map((c) => ssr(re, ssrHydrationKey(), escape(c.title), escape(new Date(c.release_date).getFullYear()))))) : escape([]), e.status == p.won ? escape(createComponent(fe$1, { disabled: true, get movie() {
    return e.correct_option;
  } })) : e.status == p.lost ? ssr(se, ssrHydrationKey(), e.selected_option ? ssr(de, ssrHydrationKey(), escape(createComponent(fe$1, { disabled: true, get movie() {
    return e.correct_option;
  } })), escape(createComponent(ye$1, { disabled: true, get movie() {
    return e.selected_option;
  } }))) : escape(createComponent(ye$1, { disabled: true, get movie() {
    return e.correct_option;
  } }))) : escape([]));
}
const I = createContext([{}, () => {
}]);
function le(e) {
  let l = makePersisted(createStore({ difficulty: u.hard }), { name: "mural_user" });
  return createComponent(I.Provider, { value: l, get children() {
    return e.children;
  } });
}
function g() {
  return useContext(I);
}
var oe = ["<div", ' class="flex space-x-2 items-center"><div class="inline-flex rounded-full text-sm text-river-bed-700 font-bold" role="group"><button type="button" class="', '">Easy Mode</button><button type="button" class="', '">Hard Mode</button></div></div>'];
let u = function(e) {
  return e[e.easy = 0] = "easy", e[e.hard = 1] = "hard", e;
}({});
function ae() {
  const [e, l] = v(), [r, a] = g();
  return ssr(oe, ssrHydrationKey(), `${r.difficulty == u.easy ? "bg-desert-sand-300 border-river-bed-700" : ""} ${r.difficulty != u.easy && e.flipped.length == 0 ? "border-river-bed-700" : ""} ${r.difficulty != u.easy && e.flipped.length != 0 ? "border-desert-sand-300 text-desert-sand-300" : ""} px-4 py-1 rounded-s-full border-r border-y-2 border-l-2`, `${r.difficulty == u.hard ? "bg-desert-sand-300 border-river-bed-700" : ""} ${r.difficulty != u.hard && e.flipped.length == 0 ? "border-river-bed-700" : ""} ${r.difficulty != u.hard && e.flipped.length != 0 ? "border-desert-sand-300 text-desert-sand-300" : ""}  px-4 py-1 rounded-e-full border-l border-y-2 border-r-2 border-river-bed-700 `);
}
var ne = ["<div", ' class="w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center border-dingley-700 border-2 hover:cursor-pointer bg-dingley-700 hover:bg-dingley-800">Share</div>'];
function ce(e) {
  return v(), g(), ssr(ne, ssrHydrationKey());
}
var ue = ["<main", ' class="w-full text-river-bed-600 font-extrabold flex flex-col space-y-4">', "</main>"], ve = ["<div", ' class="h-0.5 w-full rounded-full bg-river-bed-600"></div>'], fe = ["<div", ' class="flex flex-col space-y-2"><!--$-->', "<!--/--><!--$-->", "<!--/--></div>"];
function be() {
  const [e, l] = v(), [r, a] = g();
  return ssr(ue, ssrHydrationKey(), e.flipped.length == 0 ? escape([]) : escape([ssr(ve, ssrHydrationKey()), e.status == p.won || e.status == p.lost ? createComponent(ce, { onclick: () => {
    window.scrollTo(0, 0), open();
  } }) : [], r.difficulty == u.easy ? createComponent(Q, {}) : createComponent(ie, {}), e.status == p.won || e.status == p.lost ? " " : ssr(fe, ssrHydrationKey(), escape(createComponent(R, {})), r.difficulty == u.hard ? escape(createComponent(W, {})) : escape([]))]));
}
var pe = ["<main", ' class="text-river-bed-600 font-extrabold flex flex-col items-center justify-center space-y-4"><div class="flex flex-col items-center space-y-2"><div class="flex space-x-2 text-4xl"><div>Score:</div><div>', "</div></div><div>", "</div></div><!--$-->", "<!--/--><!--$-->", "<!--/--></main>"];
function me() {
  const [e, l] = v();
  return ssr(pe, ssrHydrationKey(), escape(e.score), escape(createComponent(ae, {})), escape(createComponent(be$1, {})), escape(createComponent(be, {})));
}
var ge = ["<div", ' class="', '"><div class="z-10 absolute top-0 left-0 right-0 bottom-0 justify-center items-center bg-black flex opacity-70"></div><div class="z-20 p-4 border-2 border-river-bed-700 absolute top-0 left-0 right-0 md:w-128 md:mx-auto m-4 rounded-lg bg-desert-sand-100 shadow-lg flex flex-col space-y-2 justify-between overflow-auto"><div id="dialog-content" class="p-4 flex flex-col space-y-2 w-full"><div id="info-dialog"><div class="flex flex-col space-y-2"><div class="flex justify-between items-center"><div class="text-3xl">Mural</div><button><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6"><path fill-rule="evenodd" d="M5.47 5.47a.75.75 0 011.06 0L12 10.94l5.47-5.47a.75.75 0 111.06 1.06L13.06 12l5.47 5.47a.75.75 0 11-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 01-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 010-1.06z" clip-rule="evenodd"></path></svg></button></div><div class="h-0.5 w-full rounded-full bg-river-bed-600"></div><div class="flex flex-col space-y-2"><div class="text-xl">What is Mural</div><div class="text-md">Mural is a daily puzzle game where you have to guess a movie poster by flipping over tiles. The puzzle refreshes daily at 12:00AM EST.</div><div>Every day has a different theme. On normal weeks the themes go by decade.</div><ul><li><strong>Monday</strong>: 2020s</li><li><strong>Tuesday</strong>: 2010s</li><li><strong>Wednesday</strong>: 2000s</li><li><strong>Thursday</strong>: 1990s</li><li><strong>Friday</strong>: 1980s</li><li><strong>Saturday</strong>: 1970s</li><li><strong>Sunday</strong>: Random</li></ul><div class="text-xl">How to Play</div><div>Flip over tiles one at a time. Each tile has a penalty. The outer tiles have the lowest penalty and the inner ones the highest.</div></div></div></div></div></div></div>'];
function xe() {
  const [e, { open: l, close: r }] = D(), [a, c] = createSignal(ge$1());
  return createEffect(() => {
    setInterval(() => {
      c(ge$1());
    }, 1e3);
  }), ssr(ge, ssrHydrationKey(), `${e.dialog_status ? "" : "hidden"} ${e.dialog_status ? "block" : ""}`);
}
var ye = ["<div", ' class="', '"><div class="z-10 absolute top-0 left-0 right-0 bottom-0 justify-center items-center bg-black flex opacity-70"></div><div class="z-20 p-4 border-2 border-river-bed-700 absolute top-0 left-0 right-0 md:w-128 md:mx-auto m-4 rounded-lg bg-desert-sand-100 shadow-lg flex flex-col space-y-2 justify-between overflow-auto"><div id="dialog-content" class="p-4 flex flex-col space-y-2 w-full"><div id="info-dialog"><div class="flex flex-col space-y-2"><div class="flex justify-between items-center"><div class="text-3xl">Hints</div><button><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6"><path fill-rule="evenodd" d="M5.47 5.47a.75.75 0 011.06 0L12 10.94l5.47-5.47a.75.75 0 111.06 1.06L13.06 12l5.47 5.47a.75.75 0 11-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 01-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 010-1.06z" clip-rule="evenodd"></path></svg></button></div><div class="h-0.5 w-full rounded-full bg-river-bed-600"></div><div class="flex flex-col space-y-2"><div>Year</div><div class="text-contessa-500 text-lg">', '</div><div>Genres</div><div class="flex space-x-2">', '</div><div>Description</div><div class="text-contessa-500 text-lg flex space-x-2">', "</div></div></div></div></div></div></div>"], h = ["<button", ' id="info-button" class=" w-full p-1 text-base text-desert-sand-100 rounded-md flex justify-center items-center border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700 ">Reveal Year (-5)</button>'], he = ["<div", ' class="rounded-full px-2 bg-contessa-500 text-el-salva-100">', "</div>"], _ = ["<button", ' id="info-button" class=" w-full p-1 text-base text-desert-sand-100 rounded-md flex justify-center items-center border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700 ">Reveal Genres (-10)</button>'], $ = ["<button", ' id="info-button" class=" w-full p-1 text-base text-desert-sand-100 rounded-md flex justify-center items-center border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700 ">Reveal Description (-25)</button>'];
function _e() {
  var _a, _b, _c;
  const [e, { open: l, close: r }] = U$1(), [a, c] = v();
  return ssr(ye, ssrHydrationKey(), `${e() ? "" : "hidden"} ${e() ? "block" : ""}`, ((_a = a.hints) == null ? void 0 : _a.year) ? escape(new Date(a.correct_option.release_date).getFullYear()) : h[0] + ssrHydrationKey() + h[1], ((_b = a.hints) == null ? void 0 : _b.genres) ? escape(a.correct_option.genres.map((f) => ssr(he, ssrHydrationKey(), escape(f.name)))) : _[0] + ssrHydrationKey() + _[1], ((_c = a.hints) == null ? void 0 : _c.description) ? escape(a.correct_option.overview) : $[0] + ssrHydrationKey() + $[1]);
}
var $e = ["<div", ' class="flex flex-col items-center justify-center"><div class="flex flex-col items-center space-y-4 p-4"><div class="flex flex-col space-y-4 w-full"><div class="rounded-lg p-4 border-2 border-river-bed-700 flex space-x-1 justify-center"><div>Play other daily games</div><a class="text-contessa-600 underline" href="https://ancgames.com">here.</a></div><div class="text-5xl flex space-x-2 items-center"><div>Mural #<!--$-->', '<!--/--></div><div id="game-version" class="font-semibold w-min h-min text-gray-600 text-xs border-2 px-1 border-river-bed-700 rounded-lg">', '</div></div><div class="flex justify-between"><div class="flex flex-col space-y-1 items-start"><div id="game-theme" class="text-contessa-500 text-4xl">', `</div><div class="text-md">Today's Theme</div></div><!--$-->`, '<!--/--></div><div class="flex flex-col space-y-1"><div class="w-full">', '</div></div><div class="h-0.5 w-full rounded-full bg-river-bed-600"></div><div class="text-3xl flex space-x-2 items-center flex-col">', "</div></div></div><!--$-->", "<!--/--><!--$-->", "<!--/--></div>"], we = ["<div", ' class="flex flex-col space-y-1 items-start"><div id="games-played" class="text-contessa-500 text-4xl">', '</div><div class="text-md">Have Played</div></div>'];
const je = () => {
  const [e, l] = v(), [r, a] = createSignal(void 0);
  return createEffect(() => {
    (e.game_key != E() || e.correct_option.id != N$1().correct_option.id) && (localStorage.removeItem("mural_game"), l(N$1()));
  }), createEffect(() => {
    pe$1().then((c) => {
      a(c);
    });
  }), [createComponent(f, { name: "viewport", content: "width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=0" }), createComponent(F, { get children() {
    return createComponent(we$1, { get children() {
      var _a;
      return ssr($e, ssrHydrationKey(), escape(e.game_key), (_a = escape({}.VITE_VERSION)) != null ? _a : "v0.1.1", escape(e.theme), r() ? ssr(we, ssrHydrationKey(), escape(r())) : escape([]), escape(createComponent(O, {})), escape(createComponent(me, {})), escape(createComponent(xe, {})), escape(createComponent(_e, {})));
    } });
  } })];
};
function Pe() {
  return createComponent(ve$1, { get children() {
    return createComponent(le, { get children() {
      return createComponent(je, {});
    } });
  } });
}

export { Pe as default };
//# sourceMappingURL=index.mjs.map
