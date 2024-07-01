globalThis._importMeta_={url:import.meta.url,env:process.env};import 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/node-fetch-native/dist/polyfill.cjs';
import { defineEventHandler, handleCacheHeaders, splitCookiesString, isEvent, createEvent, getRequestHeader, eventHandler, setHeaders, sendRedirect, proxyRequest, setResponseStatus, setResponseHeader, send, removeResponseHeader, createError, getResponseHeader, setHeader, getRequestIP, getRequestURL, getRequestWebStream, appendResponseHeader, getCookie, setCookie, createApp, createRouter as createRouter$1, toNodeListener, fetchWithEvent, lazyEventHandler } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/h3/dist/index.mjs';
import { createFetch as createFetch$1, Headers as Headers$1 } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/ofetch/dist/node.mjs';
import destr from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/destr/dist/index.mjs';
import { createCall, createFetch } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/unenv/runtime/fetch/index.mjs';
import { createHooks } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/hookable/dist/index.mjs';
import { snakeCase } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/scule/dist/index.mjs';
import { klona } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/klona/dist/index.mjs';
import defu, { defuFn } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/defu/dist/defu.mjs';
import { hash } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/ohash/dist/index.mjs';
import { parseURL, withoutBase, joinURL, getQuery, withQuery, decodePath, withLeadingSlash, withoutTrailingSlash } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/ufo/dist/index.mjs';
import { createStorage, prefixStorage } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/unstorage/dist/index.mjs';
import unstorage_47drivers_47fs from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/unstorage/drivers/fs.mjs';
import unstorage_47drivers_47fs_45lite from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/unstorage/drivers/fs-lite.mjs';
import { toRouteMatcher, createRouter } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/radix3/dist/index.mjs';
import _PlvCE0J2Zj from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/vinxi/lib/app-fetch.js';
import _Uosyeeba5I from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/vinxi/lib/app-manifest.js';
import { promises } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { dirname, resolve } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/pathe/dist/index.mjs';
import { fromJSON, crossSerializeStream, getCrossReferenceHeader } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/seroval/dist/esm/production/index.mjs';
import { CustomEventPlugin, DOMExceptionPlugin, EventPlugin, FormDataPlugin, HeadersPlugin, ReadableStreamPlugin, RequestPlugin, ResponsePlugin, URLSearchParamsPlugin, URLPlugin } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/seroval-plugins/dist/esm/production/web.mjs';
import { sharedConfig, lazy, createComponent as createComponent$1 } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/solid-js/dist/server.js';
import { provideRequestEvent } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/solid-js/web/dist/storage.js';
import { ssr, createComponent, ssrHydrationKey, NoHydration, escape, getRequestEvent, ssrAttribute, renderToStream, ssrElement, mergeProps } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/node_modules/solid-js/web/dist/server.js';

const inlineAppConfig = {};



const appConfig$1 = defuFn(inlineAppConfig);

const _inlineRuntimeConfig = {
  "app": {
    "baseURL": "/"
  },
  "nitro": {
    "routeRules": {}
  }
};
const ENV_PREFIX = "NITRO_";
const ENV_PREFIX_ALT = _inlineRuntimeConfig.nitro.envPrefix ?? process.env.NITRO_ENV_PREFIX ?? "_";
const _sharedRuntimeConfig = _deepFreeze(
  _applyEnv(klona(_inlineRuntimeConfig))
);
function useRuntimeConfig(event) {
  if (!event) {
    return _sharedRuntimeConfig;
  }
  if (event.context.nitro.runtimeConfig) {
    return event.context.nitro.runtimeConfig;
  }
  const runtimeConfig = klona(_inlineRuntimeConfig);
  _applyEnv(runtimeConfig);
  event.context.nitro.runtimeConfig = runtimeConfig;
  return runtimeConfig;
}
_deepFreeze(klona(appConfig$1));
function _getEnv(key) {
  const envKey = snakeCase(key).toUpperCase();
  return destr(
    process.env[ENV_PREFIX + envKey] ?? process.env[ENV_PREFIX_ALT + envKey]
  );
}
function _isObject(input) {
  return typeof input === "object" && !Array.isArray(input);
}
function _applyEnv(obj, parentKey = "") {
  for (const key in obj) {
    const subKey = parentKey ? `${parentKey}_${key}` : key;
    const envValue = _getEnv(subKey);
    if (_isObject(obj[key])) {
      if (_isObject(envValue)) {
        obj[key] = { ...obj[key], ...envValue };
      }
      _applyEnv(obj[key], subKey);
    } else {
      obj[key] = envValue ?? obj[key];
    }
  }
  return obj;
}
function _deepFreeze(object) {
  const propNames = Object.getOwnPropertyNames(object);
  for (const name of propNames) {
    const value = object[name];
    if (value && typeof value === "object") {
      _deepFreeze(value);
    }
  }
  return Object.freeze(object);
}
new Proxy(/* @__PURE__ */ Object.create(null), {
  get: (_, prop) => {
    console.warn(
      "Please use `useRuntimeConfig()` instead of accessing config directly."
    );
    const runtimeConfig = useRuntimeConfig();
    if (prop in runtimeConfig) {
      return runtimeConfig[prop];
    }
    return void 0;
  }
});

const serverAssets = [{"baseName":"server","dir":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/assets"}];

const assets$1 = createStorage();

for (const asset of serverAssets) {
  assets$1.mount(asset.baseName, unstorage_47drivers_47fs({ base: asset.dir }));
}

const storage = createStorage({});

storage.mount('/assets', assets$1);

storage.mount('data', unstorage_47drivers_47fs_45lite({"driver":"fsLite","base":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/.data/kv"}));
storage.mount('root', unstorage_47drivers_47fs({"driver":"fs","readOnly":true,"base":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code","ignore":["**/node_modules/**","**/.git/**"]}));
storage.mount('src', unstorage_47drivers_47fs({"driver":"fs","readOnly":true,"base":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code","ignore":["**/node_modules/**","**/.git/**"]}));
storage.mount('build', unstorage_47drivers_47fs({"driver":"fs","readOnly":false,"base":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/.vinxi","ignore":["**/node_modules/**","**/.git/**"]}));
storage.mount('cache', unstorage_47drivers_47fs({"driver":"fs","readOnly":false,"base":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/.vinxi/cache","ignore":["**/node_modules/**","**/.git/**"]}));

function useStorage(base = "") {
  return base ? prefixStorage(storage, base) : storage;
}

const defaultCacheOptions = {
  name: "_",
  base: "/cache",
  swr: true,
  maxAge: 1
};
function defineCachedFunction(fn, opts = {}) {
  opts = { ...defaultCacheOptions, ...opts };
  const pending = {};
  const group = opts.group || "nitro/functions";
  const name = opts.name || fn.name || "_";
  const integrity = opts.integrity || hash([fn, opts]);
  const validate = opts.validate || ((entry) => entry.value !== void 0);
  async function get(key, resolver, shouldInvalidateCache, event) {
    const cacheKey = [opts.base, group, name, key + ".json"].filter(Boolean).join(":").replace(/:\/$/, ":index");
    const entry = await useStorage().getItem(cacheKey) || {};
    const ttl = (opts.maxAge ?? opts.maxAge ?? 0) * 1e3;
    if (ttl) {
      entry.expires = Date.now() + ttl;
    }
    const expired = shouldInvalidateCache || entry.integrity !== integrity || ttl && Date.now() - (entry.mtime || 0) > ttl || validate(entry) === false;
    const _resolve = async () => {
      const isPending = pending[key];
      if (!isPending) {
        if (entry.value !== void 0 && (opts.staleMaxAge || 0) >= 0 && opts.swr === false) {
          entry.value = void 0;
          entry.integrity = void 0;
          entry.mtime = void 0;
          entry.expires = void 0;
        }
        pending[key] = Promise.resolve(resolver());
      }
      try {
        entry.value = await pending[key];
      } catch (error) {
        if (!isPending) {
          delete pending[key];
        }
        throw error;
      }
      if (!isPending) {
        entry.mtime = Date.now();
        entry.integrity = integrity;
        delete pending[key];
        if (validate(entry) !== false) {
          const promise = useStorage().setItem(cacheKey, entry).catch((error) => {
            console.error(`[nitro] [cache] Cache write error.`, error);
            useNitroApp().captureError(error, { event, tags: ["cache"] });
          });
          if (event && event.waitUntil) {
            event.waitUntil(promise);
          }
        }
      }
    };
    const _resolvePromise = expired ? _resolve() : Promise.resolve();
    if (entry.value === void 0) {
      await _resolvePromise;
    } else if (expired && event && event.waitUntil) {
      event.waitUntil(_resolvePromise);
    }
    if (opts.swr && validate(entry) !== false) {
      _resolvePromise.catch((error) => {
        console.error(`[nitro] [cache] SWR handler error.`, error);
        useNitroApp().captureError(error, { event, tags: ["cache"] });
      });
      return entry;
    }
    return _resolvePromise.then(() => entry);
  }
  return async (...args) => {
    const shouldBypassCache = opts.shouldBypassCache?.(...args);
    if (shouldBypassCache) {
      return fn(...args);
    }
    const key = await (opts.getKey || getKey)(...args);
    const shouldInvalidateCache = opts.shouldInvalidateCache?.(...args);
    const entry = await get(
      key,
      () => fn(...args),
      shouldInvalidateCache,
      args[0] && isEvent(args[0]) ? args[0] : void 0
    );
    let value = entry.value;
    if (opts.transform) {
      value = await opts.transform(entry, ...args) || value;
    }
    return value;
  };
}
const cachedFunction = defineCachedFunction;
function getKey(...args) {
  return args.length > 0 ? hash(args, {}) : "";
}
function escapeKey(key) {
  return String(key).replace(/\W/g, "");
}
function defineCachedEventHandler(handler, opts = defaultCacheOptions) {
  const variableHeaderNames = (opts.varies || []).filter(Boolean).map((h) => h.toLowerCase()).sort();
  const _opts = {
    ...opts,
    getKey: async (event) => {
      const customKey = await opts.getKey?.(event);
      if (customKey) {
        return escapeKey(customKey);
      }
      const _path = event.node.req.originalUrl || event.node.req.url || event.path;
      const _pathname = escapeKey(decodeURI(parseURL(_path).pathname)).slice(0, 16) || "index";
      const _hashedPath = `${_pathname}.${hash(_path)}`;
      const _headers = variableHeaderNames.map((header) => [header, event.node.req.headers[header]]).map(([name, value]) => `${escapeKey(name)}.${hash(value)}`);
      return [_hashedPath, ..._headers].join(":");
    },
    validate: (entry) => {
      if (!entry.value) {
        return false;
      }
      if (entry.value.code >= 400) {
        return false;
      }
      if (entry.value.body === void 0) {
        return false;
      }
      if (entry.value.headers.etag === "undefined" || entry.value.headers["last-modified"] === "undefined") {
        return false;
      }
      return true;
    },
    group: opts.group || "nitro/handlers",
    integrity: opts.integrity || hash([handler, opts])
  };
  const _cachedHandler = cachedFunction(
    async (incomingEvent) => {
      const variableHeaders = {};
      for (const header of variableHeaderNames) {
        variableHeaders[header] = incomingEvent.node.req.headers[header];
      }
      const reqProxy = cloneWithProxy(incomingEvent.node.req, {
        headers: variableHeaders
      });
      const resHeaders = {};
      let _resSendBody;
      const resProxy = cloneWithProxy(incomingEvent.node.res, {
        statusCode: 200,
        writableEnded: false,
        writableFinished: false,
        headersSent: false,
        closed: false,
        getHeader(name) {
          return resHeaders[name];
        },
        setHeader(name, value) {
          resHeaders[name] = value;
          return this;
        },
        getHeaderNames() {
          return Object.keys(resHeaders);
        },
        hasHeader(name) {
          return name in resHeaders;
        },
        removeHeader(name) {
          delete resHeaders[name];
        },
        getHeaders() {
          return resHeaders;
        },
        end(chunk, arg2, arg3) {
          if (typeof chunk === "string") {
            _resSendBody = chunk;
          }
          if (typeof arg2 === "function") {
            arg2();
          }
          if (typeof arg3 === "function") {
            arg3();
          }
          return this;
        },
        write(chunk, arg2, arg3) {
          if (typeof chunk === "string") {
            _resSendBody = chunk;
          }
          if (typeof arg2 === "function") {
            arg2();
          }
          if (typeof arg3 === "function") {
            arg3();
          }
          return this;
        },
        writeHead(statusCode, headers2) {
          this.statusCode = statusCode;
          if (headers2) {
            for (const header in headers2) {
              this.setHeader(header, headers2[header]);
            }
          }
          return this;
        }
      });
      const event = createEvent(reqProxy, resProxy);
      event.context = incomingEvent.context;
      const body = await handler(event) || _resSendBody;
      const headers = event.node.res.getHeaders();
      headers.etag = String(
        headers.Etag || headers.etag || `W/"${hash(body)}"`
      );
      headers["last-modified"] = String(
        headers["Last-Modified"] || headers["last-modified"] || (/* @__PURE__ */ new Date()).toUTCString()
      );
      const cacheControl = [];
      if (opts.swr) {
        if (opts.maxAge) {
          cacheControl.push(`s-maxage=${opts.maxAge}`);
        }
        if (opts.staleMaxAge) {
          cacheControl.push(`stale-while-revalidate=${opts.staleMaxAge}`);
        } else {
          cacheControl.push("stale-while-revalidate");
        }
      } else if (opts.maxAge) {
        cacheControl.push(`max-age=${opts.maxAge}`);
      }
      if (cacheControl.length > 0) {
        headers["cache-control"] = cacheControl.join(", ");
      }
      const cacheEntry = {
        code: event.node.res.statusCode,
        headers,
        body
      };
      return cacheEntry;
    },
    _opts
  );
  return defineEventHandler(async (event) => {
    if (opts.headersOnly) {
      if (handleCacheHeaders(event, { maxAge: opts.maxAge })) {
        return;
      }
      return handler(event);
    }
    const response = await _cachedHandler(event);
    if (event.node.res.headersSent || event.node.res.writableEnded) {
      return response.body;
    }
    if (handleCacheHeaders(event, {
      modifiedTime: new Date(response.headers["last-modified"]),
      etag: response.headers.etag,
      maxAge: opts.maxAge
    })) {
      return;
    }
    event.node.res.statusCode = response.code;
    for (const name in response.headers) {
      const value = response.headers[name];
      if (name === "set-cookie") {
        event.node.res.appendHeader(
          name,
          splitCookiesString(value)
        );
      } else {
        event.node.res.setHeader(name, value);
      }
    }
    return response.body;
  });
}
function cloneWithProxy(obj, overrides) {
  return new Proxy(obj, {
    get(target, property, receiver) {
      if (property in overrides) {
        return overrides[property];
      }
      return Reflect.get(target, property, receiver);
    },
    set(target, property, value, receiver) {
      if (property in overrides) {
        overrides[property] = value;
        return true;
      }
      return Reflect.set(target, property, value, receiver);
    }
  });
}
const cachedEventHandler = defineCachedEventHandler;

function hasReqHeader(event, name, includes) {
  const value = getRequestHeader(event, name);
  return value && typeof value === "string" && value.toLowerCase().includes(includes);
}
function isJsonRequest(event) {
  if (hasReqHeader(event, "accept", "text/html")) {
    return false;
  }
  return hasReqHeader(event, "accept", "application/json") || hasReqHeader(event, "user-agent", "curl/") || hasReqHeader(event, "user-agent", "httpie/") || hasReqHeader(event, "sec-fetch-mode", "cors") || event.path.startsWith("/api/") || event.path.endsWith(".json");
}
function normalizeError(error) {
  const cwd = typeof process.cwd === "function" ? process.cwd() : "/";
  const stack = (error.stack || "").split("\n").splice(1).filter((line) => line.includes("at ")).map((line) => {
    const text = line.replace(cwd + "/", "./").replace("webpack:/", "").replace("file://", "").trim();
    return {
      text,
      internal: line.includes("node_modules") && !line.includes(".cache") || line.includes("internal") || line.includes("new Promise")
    };
  });
  const statusCode = error.statusCode || 500;
  const statusMessage = error.statusMessage ?? (statusCode === 404 ? "Not Found" : "");
  const message = error.message || error.toString();
  return {
    stack,
    statusCode,
    statusMessage,
    message
  };
}
function _captureError(error, type) {
  console.error(`[nitro] [${type}]`, error);
  useNitroApp().captureError(error, { tags: [type] });
}
function trapUnhandledNodeErrors() {
  process.on(
    "unhandledRejection",
    (error) => _captureError(error, "unhandledRejection")
  );
  process.on(
    "uncaughtException",
    (error) => _captureError(error, "uncaughtException")
  );
}
function joinHeaders(value) {
  return Array.isArray(value) ? value.join(", ") : String(value);
}
function normalizeFetchResponse(response) {
  if (!response.headers.has("set-cookie")) {
    return response;
  }
  return new Response(response.body, {
    status: response.status,
    statusText: response.statusText,
    headers: normalizeCookieHeaders(response.headers)
  });
}
function normalizeCookieHeader(header = "") {
  return splitCookiesString(joinHeaders(header));
}
function normalizeCookieHeaders(headers) {
  const outgoingHeaders = new Headers();
  for (const [name, header] of headers) {
    if (name === "set-cookie") {
      for (const cookie of normalizeCookieHeader(header)) {
        outgoingHeaders.append("set-cookie", cookie);
      }
    } else {
      outgoingHeaders.set(name, joinHeaders(header));
    }
  }
  return outgoingHeaders;
}

const config = useRuntimeConfig();
const _routeRulesMatcher = toRouteMatcher(
  createRouter({ routes: config.nitro.routeRules })
);
function createRouteRulesHandler(ctx) {
  return eventHandler((event) => {
    const routeRules = getRouteRules(event);
    if (routeRules.headers) {
      setHeaders(event, routeRules.headers);
    }
    if (routeRules.redirect) {
      return sendRedirect(
        event,
        routeRules.redirect.to,
        routeRules.redirect.statusCode
      );
    }
    if (routeRules.proxy) {
      let target = routeRules.proxy.to;
      if (target.endsWith("/**")) {
        let targetPath = event.path;
        const strpBase = routeRules.proxy._proxyStripBase;
        if (strpBase) {
          targetPath = withoutBase(targetPath, strpBase);
        }
        target = joinURL(target.slice(0, -3), targetPath);
      } else if (event.path.includes("?")) {
        const query = getQuery(event.path);
        target = withQuery(target, query);
      }
      return proxyRequest(event, target, {
        fetch: ctx.localFetch,
        ...routeRules.proxy
      });
    }
  });
}
function getRouteRules(event) {
  event.context._nitro = event.context._nitro || {};
  if (!event.context._nitro.routeRules) {
    event.context._nitro.routeRules = getRouteRulesForPath(
      withoutBase(event.path.split("?")[0], useRuntimeConfig().app.baseURL)
    );
  }
  return event.context._nitro.routeRules;
}
function getRouteRulesForPath(path) {
  return defu({}, ..._routeRulesMatcher.matchAll(path).reverse());
}

const appConfig = {"name":"vinxi","routers":[{"name":"public","mode":"static","dir":"./public","base":"/","root":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code","order":0,"outDir":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/.vinxi/build/public"},{"name":"ssr","mode":"handler","handler":"src/entry-server.tsx","extensions":["js","jsx","ts","tsx"],"target":"server","root":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code","base":"/","outDir":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/.vinxi/build/ssr","order":1},{"name":"client","mode":"build","handler":"src/entry-client.tsx","extensions":["js","jsx","ts","tsx"],"target":"browser","base":"/_build","root":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code","outDir":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/.vinxi/build/client","order":2},{"name":"server-fns","mode":"handler","base":"/_server","handler":"node_modules/@solidjs/start/config/server-handler.js","target":"server","root":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code","outDir":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/.vinxi/build/server-fns","order":3}],"server":{"compressPublicAssets":{"brotli":true},"baseURL":"/","prerender":{}},"root":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code"};
				const buildManifest = {"ssr":{"_game-board-b1526277.js":{"file":"assets/game-board-b1526277.js"},"_index-d741c3e0.js":{"file":"assets/index-d741c3e0.js"},"src/routes/[...404].tsx?pick=default&pick=$css":{"file":"_...404_.js","imports":["_index-d741c3e0.js"],"isDynamicEntry":true,"isEntry":true,"src":"src/routes/[...404].tsx?pick=default&pick=$css"},"src/routes/index.tsx?pick=default&pick=$css":{"file":"index.js","imports":["_game-board-b1526277.js","_index-d741c3e0.js"],"isDynamicEntry":true,"isEntry":true,"src":"src/routes/index.tsx?pick=default&pick=$css"},"src/routes/share.tsx?pick=default&pick=$css":{"file":"share.js","imports":["_index-d741c3e0.js","_game-board-b1526277.js"],"isDynamicEntry":true,"isEntry":true,"src":"src/routes/share.tsx?pick=default&pick=$css"},"virtual:#vinxi/handler/ssr":{"dynamicImports":["src/routes/[...404].tsx?pick=default&pick=$css","src/routes/[...404].tsx?pick=default&pick=$css","src/routes/index.tsx?pick=default&pick=$css","src/routes/index.tsx?pick=default&pick=$css","src/routes/share.tsx?pick=default&pick=$css","src/routes/share.tsx?pick=default&pick=$css"],"file":"ssr.js","isEntry":true,"src":"virtual:#vinxi/handler/ssr"}},"client":{"\u0000virtual:#vinxi/handler/client.css":{"file":"assets/client-653d473a.css","src":"\u0000virtual:#vinxi/handler/client.css"},"_game-board-36969377.js":{"file":"assets/game-board-36969377.js","imports":["_index-16c4c740.js"]},"_index-16c4c740.js":{"file":"assets/index-16c4c740.js"},"_routing-8c4bfa0f.js":{"file":"assets/routing-8c4bfa0f.js"},"node_modules/shikiji/dist/onig.wasm":{"file":"assets/onig-fd885c2d.wasm","src":"node_modules/shikiji/dist/onig.wasm"},"src/routes/[...404].tsx?pick=default&pick=$css":{"file":"assets/_...404_-0c4dce42.js","imports":["_index-16c4c740.js"],"isDynamicEntry":true,"isEntry":true,"src":"src/routes/[...404].tsx?pick=default&pick=$css"},"src/routes/index.tsx?pick=default&pick=$css":{"file":"assets/index-905008f7.js","imports":["_index-16c4c740.js","_game-board-36969377.js"],"isDynamicEntry":true,"isEntry":true,"src":"src/routes/index.tsx?pick=default&pick=$css"},"src/routes/share.tsx?pick=default&pick=$css":{"file":"assets/share-eefaf08c.js","imports":["_index-16c4c740.js","_game-board-36969377.js","_routing-8c4bfa0f.js"],"isDynamicEntry":true,"isEntry":true,"src":"src/routes/share.tsx?pick=default&pick=$css"},"virtual:#vinxi/handler/client":{"css":["assets/client-653d473a.css"],"dynamicImports":["src/routes/[...404].tsx?pick=default&pick=$css","src/routes/index.tsx?pick=default&pick=$css","src/routes/share.tsx?pick=default&pick=$css"],"file":"assets/client-27e5df91.js","imports":["_index-16c4c740.js","_routing-8c4bfa0f.js"],"isEntry":true,"src":"virtual:#vinxi/handler/client"}},"server-fns":{"virtual:#vinxi/handler/server-fns":{"file":"entry.js","isEntry":true,"src":"virtual:#vinxi/handler/server-fns"}}};

				const routeManifest = {"ssr":{},"client":{}};

        function createProdApp(appConfig) {
          return {
            config: { ...appConfig, buildManifest, routeManifest },
            getRouter(name) {
              return appConfig.routers.find(router => router.name === name)
            }
          }
        }

        function plugin(app) {
          const prodApp = createProdApp(appConfig);
          globalThis.app = prodApp;
        }

const chunks = {};
			 



			 function app() {
				 globalThis.$$chunks = chunks;
			 }

const plugins = [
  plugin,
_PlvCE0J2Zj,
_Uosyeeba5I,
app
];

function defineNitroErrorHandler(handler) {
  return handler;
}
const errorHandler = defineNitroErrorHandler(
  function defaultNitroErrorHandler(error, event) {
    const { stack, statusCode, statusMessage, message } = normalizeError(error);
    const errorObject = {
      url: event.path || "",
      statusCode,
      statusMessage,
      message,
      stack: void 0
    };
    if (error.unhandled || error.fatal) {
      const tags = [
        "[nitro]",
        "[request error]",
        error.unhandled && "[unhandled]",
        error.fatal && "[fatal]"
      ].filter(Boolean).join(" ");
      console.error(
        tags,
        error.message + "\n" + stack.map((l) => "  " + l.text).join("  \n")
      );
    }
    setResponseStatus(event, statusCode, statusMessage);
    if (isJsonRequest(event)) {
      setResponseHeader(event, "Content-Type", "application/json");
      return send(event, JSON.stringify(errorObject));
    } else {
      setResponseHeader(event, "Content-Type", "text/html");
      return send(event, renderHTMLError(errorObject));
    }
  }
);
function renderHTMLError(error) {
  const statusCode = error.statusCode || 500;
  const statusMessage = error.statusMessage || "Request Error";
  return `<!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>${statusCode} ${statusMessage}</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico/css/pico.min.css">
  </head>
  <body>
    <main class="container">
      <dialog open>
        <article>
          <header>
            <h2>${statusCode} ${statusMessage}</h2>
          </header>
          <code>
            ${error.message}<br><br>
            ${"\n" + (error.stack || []).map((i) => `&nbsp;&nbsp;${i}`).join("<br>")}
          </code>
          <footer>
            <a href="/" onclick="event.preventDefault();history.back();">Go Back</a>
          </footer>
        </article>
      </dialog>
    </main>
  </body>
</html>
`;
}

const assets = {
  "/_build/manifest.json": {
    "type": "application/json",
    "etag": "\"76d-T1thYahEjsEyVn3AXtH6pplpaVM\"",
    "mtime": "2024-07-01T16:54:32.846Z",
    "size": 1901,
    "path": "../../.output/public/_build/manifest.json"
  },
  "/_build/manifest.json.br": {
    "type": "application/json",
    "encoding": "br",
    "etag": "\"192-XEUibS3ORX5FIgGyh5vgiE5OSj8\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 402,
    "path": "../../.output/public/_build/manifest.json.br"
  },
  "/_build/manifest.json.gz": {
    "type": "application/json",
    "encoding": "gzip",
    "etag": "\"1b4-9DJ9g4BnVC7nzig9ERozZrci/5M\"",
    "mtime": "2024-07-01T16:54:32.884Z",
    "size": 436,
    "path": "../../.output/public/_build/manifest.json.gz"
  },
  "/_build/server-functions-manifest.json": {
    "type": "application/json",
    "etag": "\"19-U+evudgPW1yE9kGumdxd/vtvk2s\"",
    "mtime": "2024-07-01T16:54:32.846Z",
    "size": 25,
    "path": "../../.output/public/_build/server-functions-manifest.json"
  },
  "/assets/game-board-b1526277.js": {
    "type": "application/javascript",
    "etag": "\"7e18a-mj02H9yVvOnf7EBBEYz8XERZSE4\"",
    "mtime": "2024-07-01T16:54:32.843Z",
    "size": 516490,
    "path": "../../.output/public/assets/game-board-b1526277.js"
  },
  "/assets/game-board-b1526277.js.br": {
    "type": "application/javascript",
    "encoding": "br",
    "etag": "\"24b2e-7bmwuZzLplk4HnaE10ZJRly4TI4\"",
    "mtime": "2024-07-01T16:54:33.573Z",
    "size": 150318,
    "path": "../../.output/public/assets/game-board-b1526277.js.br"
  },
  "/assets/game-board-b1526277.js.gz": {
    "type": "application/javascript",
    "encoding": "gzip",
    "etag": "\"2e206-WOhRx7RBJEHjGiSSZHuX9d7ngsY\"",
    "mtime": "2024-07-01T16:54:32.900Z",
    "size": 188934,
    "path": "../../.output/public/assets/game-board-b1526277.js.gz"
  },
  "/assets/index-d741c3e0.js": {
    "type": "application/javascript",
    "etag": "\"1c1-whMh63tl/e/4wSS82fWoY9snkOI\"",
    "mtime": "2024-07-01T16:54:32.842Z",
    "size": 449,
    "path": "../../.output/public/assets/index-d741c3e0.js"
  },
  "/_build/assets/_...404_-0c4dce42.js": {
    "type": "application/javascript",
    "etag": "\"129-HsiPZRXacJ43luCix8TgsxFpd/o\"",
    "mtime": "2024-07-01T16:54:32.847Z",
    "size": 297,
    "path": "../../.output/public/_build/assets/_...404_-0c4dce42.js"
  },
  "/_build/assets/client-27e5df91.js": {
    "type": "application/javascript",
    "etag": "\"30ec-IZr4z/jdK6M2KzEvFjF7/2SjiMM\"",
    "mtime": "2024-07-01T16:54:32.846Z",
    "size": 12524,
    "path": "../../.output/public/_build/assets/client-27e5df91.js"
  },
  "/_build/assets/client-27e5df91.js.br": {
    "type": "application/javascript",
    "encoding": "br",
    "etag": "\"1182-u5oCiu3n2Nhto1AWzxsyPY77Lr4\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 4482,
    "path": "../../.output/public/_build/assets/client-27e5df91.js.br"
  },
  "/_build/assets/client-27e5df91.js.gz": {
    "type": "application/javascript",
    "encoding": "gzip",
    "etag": "\"1383-pABJC6K4T5ZXx/EMjvcrrp4jwV0\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 4995,
    "path": "../../.output/public/_build/assets/client-27e5df91.js.gz"
  },
  "/_build/assets/client-653d473a.css": {
    "type": "text/css; charset=utf-8",
    "etag": "\"3510-csJCIj5iwj6kyGsgmCslL8RTVFU\"",
    "mtime": "2024-07-01T16:54:32.846Z",
    "size": 13584,
    "path": "../../.output/public/_build/assets/client-653d473a.css"
  },
  "/_build/assets/client-653d473a.css.br": {
    "type": "text/css; charset=utf-8",
    "encoding": "br",
    "etag": "\"a79-QScIpsBgH1Rrv6Xb56IxhjrRKAQ\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 2681,
    "path": "../../.output/public/_build/assets/client-653d473a.css.br"
  },
  "/_build/assets/client-653d473a.css.gz": {
    "type": "text/css; charset=utf-8",
    "encoding": "gzip",
    "etag": "\"c50-odkG8hibcJi7s66hGlHQQg5+6rU\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 3152,
    "path": "../../.output/public/_build/assets/client-653d473a.css.gz"
  },
  "/_build/assets/game-board-36969377.js": {
    "type": "application/javascript",
    "etag": "\"81f75-oD0zxhx32l1KYMSVIfGhIFznt3M\"",
    "mtime": "2024-07-01T16:54:32.848Z",
    "size": 532341,
    "path": "../../.output/public/_build/assets/game-board-36969377.js"
  },
  "/_build/assets/game-board-36969377.js.br": {
    "type": "application/javascript",
    "encoding": "br",
    "etag": "\"260ca-gLeXxXNAnDG7icuRgEZHsOuZ/YY\"",
    "mtime": "2024-07-01T16:54:33.639Z",
    "size": 155850,
    "path": "../../.output/public/_build/assets/game-board-36969377.js.br"
  },
  "/_build/assets/game-board-36969377.js.gz": {
    "type": "application/javascript",
    "encoding": "gzip",
    "etag": "\"2fb27-LBgpV6f5RTAu2OUWcxKspy0HXUY\"",
    "mtime": "2024-07-01T16:54:32.901Z",
    "size": 195367,
    "path": "../../.output/public/_build/assets/game-board-36969377.js.gz"
  },
  "/_build/assets/index-16c4c740.js": {
    "type": "application/javascript",
    "etag": "\"4f16-nTIYYOLmNnGJjV95ufejjBdWn0g\"",
    "mtime": "2024-07-01T16:54:32.846Z",
    "size": 20246,
    "path": "../../.output/public/_build/assets/index-16c4c740.js"
  },
  "/_build/assets/index-16c4c740.js.br": {
    "type": "application/javascript",
    "encoding": "br",
    "etag": "\"1c93-YA54zHxj3A7SakUyIgTH50lv2gc\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 7315,
    "path": "../../.output/public/_build/assets/index-16c4c740.js.br"
  },
  "/_build/assets/index-16c4c740.js.gz": {
    "type": "application/javascript",
    "encoding": "gzip",
    "etag": "\"1f7c-8IYOVqsXzoMye70meZDIRdlyLF4\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 8060,
    "path": "../../.output/public/_build/assets/index-16c4c740.js.gz"
  },
  "/_build/assets/index-905008f7.js": {
    "type": "application/javascript",
    "etag": "\"41fa-xumoj7mH7GyXQPR83nl3fyFvEf8\"",
    "mtime": "2024-07-01T16:54:32.847Z",
    "size": 16890,
    "path": "../../.output/public/_build/assets/index-905008f7.js"
  },
  "/_build/assets/index-905008f7.js.br": {
    "type": "application/javascript",
    "encoding": "br",
    "etag": "\"105d-vZPtVgsMeTIHczUHe6TebfCzJtk\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 4189,
    "path": "../../.output/public/_build/assets/index-905008f7.js.br"
  },
  "/_build/assets/index-905008f7.js.gz": {
    "type": "application/javascript",
    "encoding": "gzip",
    "etag": "\"12b9-p02jtNeubaeah6fEBqQ7Kc0s8pQ\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 4793,
    "path": "../../.output/public/_build/assets/index-905008f7.js.gz"
  },
  "/_build/assets/onig-fd885c2d.wasm": {
    "type": "application/wasm",
    "etag": "\"71eb2-sWRKnnRfEwkGBTHc1IeyNBzS478\"",
    "mtime": "2024-07-01T16:54:32.848Z",
    "size": 466610,
    "path": "../../.output/public/_build/assets/onig-fd885c2d.wasm"
  },
  "/_build/assets/routing-8c4bfa0f.js": {
    "type": "application/javascript",
    "etag": "\"3047-gXnIziZ8e/b0/LSbHHWvxmq4LUo\"",
    "mtime": "2024-07-01T16:54:32.847Z",
    "size": 12359,
    "path": "../../.output/public/_build/assets/routing-8c4bfa0f.js"
  },
  "/_build/assets/routing-8c4bfa0f.js.br": {
    "type": "application/javascript",
    "encoding": "br",
    "etag": "\"12ca-Cj4NHBH/tnCY+842oegYzBtAkM4\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 4810,
    "path": "../../.output/public/_build/assets/routing-8c4bfa0f.js.br"
  },
  "/_build/assets/routing-8c4bfa0f.js.gz": {
    "type": "application/javascript",
    "encoding": "gzip",
    "etag": "\"1493-/fZtjEcICWrJ62ex6JcLxwophsc\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 5267,
    "path": "../../.output/public/_build/assets/routing-8c4bfa0f.js.gz"
  },
  "/_build/assets/share-eefaf08c.js": {
    "type": "application/javascript",
    "etag": "\"167e-1ABQ1Ob3000GeDD8cG6q7k+L5is\"",
    "mtime": "2024-07-01T16:54:32.847Z",
    "size": 5758,
    "path": "../../.output/public/_build/assets/share-eefaf08c.js"
  },
  "/_build/assets/share-eefaf08c.js.br": {
    "type": "application/javascript",
    "encoding": "br",
    "etag": "\"7fd-CMpdyNOZYonx3/uTwCcBffVLFc8\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 2045,
    "path": "../../.output/public/_build/assets/share-eefaf08c.js.br"
  },
  "/_build/assets/share-eefaf08c.js.gz": {
    "type": "application/javascript",
    "encoding": "gzip",
    "etag": "\"940-O3yd5KgQyBbYhISPvPT13iCJ5W8\"",
    "mtime": "2024-07-01T16:54:32.886Z",
    "size": 2368,
    "path": "../../.output/public/_build/assets/share-eefaf08c.js.gz"
  }
};

function readAsset (id) {
  const serverDir = dirname(fileURLToPath(globalThis._importMeta_.url));
  return promises.readFile(resolve(serverDir, assets[id].path))
}

const publicAssetBases = {};

function isPublicAssetURL(id = '') {
  if (assets[id]) {
    return true
  }
  for (const base in publicAssetBases) {
    if (id.startsWith(base)) { return true }
  }
  return false
}

function getAsset (id) {
  return assets[id]
}

const METHODS = /* @__PURE__ */ new Set(["HEAD", "GET"]);
const EncodingMap = { gzip: ".gz", br: ".br" };
const _f4b49z = eventHandler((event) => {
  if (event.method && !METHODS.has(event.method)) {
    return;
  }
  let id = decodePath(
    withLeadingSlash(withoutTrailingSlash(parseURL(event.path).pathname))
  );
  let asset;
  const encodingHeader = String(
    getRequestHeader(event, "accept-encoding") || ""
  );
  const encodings = [
    ...encodingHeader.split(",").map((e) => EncodingMap[e.trim()]).filter(Boolean).sort(),
    ""
  ];
  if (encodings.length > 1) {
    setResponseHeader(event, "Vary", "Accept-Encoding");
  }
  for (const encoding of encodings) {
    for (const _id of [id + encoding, joinURL(id, "index.html" + encoding)]) {
      const _asset = getAsset(_id);
      if (_asset) {
        asset = _asset;
        id = _id;
        break;
      }
    }
  }
  if (!asset) {
    if (isPublicAssetURL(id)) {
      removeResponseHeader(event, "Cache-Control");
      throw createError({
        statusMessage: "Cannot find static asset " + id,
        statusCode: 404
      });
    }
    return;
  }
  const ifNotMatch = getRequestHeader(event, "if-none-match") === asset.etag;
  if (ifNotMatch) {
    setResponseStatus(event, 304, "Not Modified");
    return "";
  }
  const ifModifiedSinceH = getRequestHeader(event, "if-modified-since");
  const mtimeDate = new Date(asset.mtime);
  if (ifModifiedSinceH && asset.mtime && new Date(ifModifiedSinceH) >= mtimeDate) {
    setResponseStatus(event, 304, "Not Modified");
    return "";
  }
  if (asset.type && !getResponseHeader(event, "Content-Type")) {
    setResponseHeader(event, "Content-Type", asset.type);
  }
  if (asset.etag && !getResponseHeader(event, "ETag")) {
    setResponseHeader(event, "ETag", asset.etag);
  }
  if (asset.mtime && !getResponseHeader(event, "Last-Modified")) {
    setResponseHeader(event, "Last-Modified", mtimeDate.toUTCString());
  }
  if (asset.encoding && !getResponseHeader(event, "Content-Encoding")) {
    setResponseHeader(event, "Content-Encoding", asset.encoding);
  }
  if (asset.size > 0 && !getResponseHeader(event, "Content-Length")) {
    setResponseHeader(event, "Content-Length", asset.size);
  }
  return readAsset(id);
});

var __defProp = Object.defineProperty;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __publicField = (obj, key, value) => {
  __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
  return value;
};
const d = "Invariant Violation", { setPrototypeOf: N = function(e, t) {
  return e.__proto__ = t, e;
} } = Object;
class p extends Error {
  constructor(t = d) {
    super(typeof t == "number" ? `${d}: ${t} (see https://github.com/apollographql/invariant-packages)` : t);
    __publicField(this, "framesToPop", 1);
    __publicField(this, "name", d);
    N(this, p.prototype);
  }
}
function m(e, t) {
  if (!e)
    throw new p(t);
}
function z(e) {
  let t;
  const r = getRequestURL(e), s = { duplex: "half", method: e.method, headers: e.headers };
  return e.node.req.body instanceof ArrayBuffer ? new Request(r, { ...s, body: e.node.req.body }) : new Request(r, { ...s, get body() {
    return t || (t = getRequestWebStream(e), t);
  } });
}
function A$1(e) {
  var _a;
  return (_a = e.web) != null ? _a : e.web = { request: z(e), url: getRequestURL(e) }, e.web.request;
}
const O = Symbol("h3Event"), f$1 = Symbol("fetchEvent"), J$1 = { get(e, t) {
  var _a;
  return t === f$1 ? e : (_a = e[t]) != null ? _a : e[O][t];
} };
function M(e) {
  const t = A$1(e);
  return new Proxy({ request: t, clientAddress: getRequestIP(e), locals: {}, [O]: e }, J$1);
}
function _(e) {
  if (!e[f$1]) {
    const t = M(e);
    e[f$1] = t;
  }
  return e[f$1];
}
function j(e) {
  const r = e.length.toString(16), s = "00000000".substring(0, 8 - r.length) + r;
  return new TextEncoder().encode(`;0x${s};${e}`);
}
function B(e, t) {
  return new ReadableStream({ start(r) {
    crossSerializeStream(t, { scopeId: e, plugins: [CustomEventPlugin, DOMExceptionPlugin, EventPlugin, FormDataPlugin, HeadersPlugin, ReadableStreamPlugin, RequestPlugin, ResponsePlugin, URLSearchParamsPlugin, URLPlugin], onSerialize(s, a) {
      r.enqueue(j(a ? `(${getCrossReferenceHeader(e)},${s})` : s));
    }, onDone() {
      r.close();
    }, onError(s) {
      r.error(s);
    } });
  } });
}
async function V$1(e) {
  m(e.method === "POST", `Invalid method ${e.method}. Expected POST.`);
  const t = _(e), r = t.request, s = r.headers.get("x-server-id"), a = r.headers.get("x-server-instance"), i = new URL(r.url);
  let c, u;
  if (s)
    m(typeof s == "string", "Invalid server function"), [c, u] = s.split("#");
  else if (c = i.searchParams.get("id"), u = i.searchParams.get("name"), !c || !u)
    throw new Error("Invalid request");
  const T = (await globalThis.MANIFEST["server-fns"].chunks[c].import())[u];
  let o = [];
  if (!a) {
    const n = i.searchParams.get("args");
    n && JSON.parse(n).forEach((l) => o.push(l));
  }
  const h = r.headers.get("content-type");
  h.startsWith("multipart/form-data") || h.startsWith("application/x-www-form-urlencoded") ? o.push(await r.formData()) : o = fromJSON(await r.json(), { plugins: [CustomEventPlugin, DOMExceptionPlugin, EventPlugin, FormDataPlugin, HeadersPlugin, ReadableStreamPlugin, RequestPlugin, ResponsePlugin, URLSearchParamsPlugin, URLPlugin] });
  try {
    const n = await provideRequestEvent(t, () => (sharedConfig.context = { event: t }, T(...o)));
    if (!a) {
      const l = n instanceof Error, H = new URL(r.headers.get("referer"));
      return new Response(null, { status: 302, headers: { Location: H.toString(), ...n ? { "Set-Cookie": `flash=${JSON.stringify({ url: i.pathname + encodeURIComponent(i.search), result: l ? n.message : n, error: l, input: [...o.slice(0, -1), [...o[o.length - 1].entries()]] })}; Secure; HttpOnly;` } : {} } });
    }
    return typeof n == "string" ? new Response(n) : (setHeader(e, "content-type", "text/javascript"), B(a, n));
  } catch (n) {
    return n instanceof Response && n.status === 302 ? new Response(null, { status: a ? 204 : 302, headers: { Location: n.headers.get("Location") } }) : n;
  }
}
const Z$1 = eventHandler(V$1);

function J(e) {
  let n;
  const t = getRequestURL(e), s = { duplex: "half", method: e.method, headers: e.headers };
  return e.node.req.body instanceof ArrayBuffer ? new Request(t, { ...s, body: e.node.req.body }) : new Request(t, { ...s, get body() {
    return n || (n = getRequestWebStream(e), n);
  } });
}
function K(e) {
  var _a;
  return (_a = e.web) != null ? _a : e.web = { request: J(e), url: getRequestURL(e) }, e.web.request;
}
const q = Symbol("h3Event"), f = Symbol("fetchEvent"), Y = { get(e, n) {
  var _a;
  return n === f ? e : (_a = e[n]) != null ? _a : e[q][n];
} };
function G(e) {
  const n = K(e);
  return new Proxy({ request: n, clientAddress: getRequestIP(e), locals: {}, [q]: e }, Y);
}
function Q(e) {
  if (!e[f]) {
    const n = G(e);
    e[f] = n;
  }
  return e[f];
}
var V = " ";
const X = { style: (e) => ssrElement("style", e.attrs, () => escape(e.children), true), link: (e) => ssrElement("link", e.attrs, void 0, true), script: (e) => e.attrs.src ? ssrElement("script", mergeProps(() => e.attrs, { get id() {
  return e.key;
} }), () => ssr(V), true) : null };
function C(e) {
  let { tag: n, attrs: { key: t, ...s } = { key: void 0 }, children: r } = e;
  return X[n]({ attrs: s, key: t, children: r });
}
var Z = ["<script", ">", "<\/script>"], ee = ["<script", ' type="module"', "><\/script>"];
const te = ssr("<!DOCTYPE html>");
function ne(e) {
  const n = getRequestEvent();
  return createComponent(NoHydration, { get children() {
    return [te, createComponent(e.document, { get assets() {
      return n.assets.map((t) => C(t));
    }, get scripts() {
      return [ssr(Z, ssrHydrationKey(), `window.manifest = ${JSON.stringify(n.manifest)}`), ssr(ee, ssrHydrationKey(), ssrAttribute("src", escape(globalThis.MANIFEST.client.inputs[globalThis.MANIFEST.client.handler].output.path, true), false))];
    } })];
  } });
}
function re(e, n, t, s = "default") {
  return lazy(async () => {
    var _a;
    {
      const o = (await e.import())[s], i = (await ((_a = n.inputs) == null ? void 0 : _a[e.src].assets())).filter((l) => l.tag === "style" || l.attrs.rel === "stylesheet");
      return { default: (l) => [...i.map((u) => C(u)), createComponent$1(o, l)] };
    }
  });
}
const T = [{ type: "page", $component: { src: "src/routes/[...404].tsx?pick=default&pick=$css", build: () => import('./chunks/build/_...404_.mjs'), import: () => import('./chunks/build/_...404_.mjs') }, path: "/*404", filePath: "/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/src/routes/[...404].tsx" }, { type: "page", $component: { src: "src/routes/index.tsx?pick=default&pick=$css", build: () => import('./chunks/build/index.mjs'), import: () => import('./chunks/build/index.mjs') }, path: "/", filePath: "/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/src/routes/index.tsx" }, { type: "page", $component: { src: "src/routes/share.tsx?pick=default&pick=$css", build: () => import('./chunks/build/share.mjs'), import: () => import('./chunks/build/share.mjs') }, path: "/share", filePath: "/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/proj-ancgames/mural/code/src/routes/share.tsx" }], ae = ce(T.filter((e) => e.type === "page")), oe = le(T.filter((e) => e.type === "api"));
function ie(e, n) {
  const t = e.replace("/", "").split("/").filter(Boolean);
  e:
    for (const s of oe) {
      const r = s.matchSegments;
      if (t.length < r.length || !s.wildcard && t.length > r.length)
        continue;
      for (let i = 0; i < r.length; i++) {
        const c = r[i];
        if (c && t[i] !== c)
          continue e;
      }
      const o = s[`$${n}`];
      if (o === "skip" || o === void 0)
        return;
      const a = {};
      for (const { type: i, name: c, index: l } of s.params)
        i === ":" ? a[c] = t[l] : a[c] = t.slice(l).join("/");
      return { handler: o, params: a };
    }
}
function ce(e) {
  function n(t, s, r, o) {
    const a = Object.values(t).find((i) => r.startsWith(i.id + "/"));
    return a ? (n(a.children || (a.children = []), s, r.slice(a.id.length)), t) : (t.push({ ...s, id: r, path: r.replace(/\/\([^)/]+\)/g, "") }), t);
  }
  return e.sort((t, s) => t.path.length - s.path.length).reduce((t, s) => n(t, s, s.path, s.path), []);
}
function le(e) {
  return e.flatMap((n) => A(n.path).map((s) => ({ ...n, path: s }))).map(ue).sort((n, t) => t.score - n.score);
}
function A(e) {
  let n = /(\/?\:[^\/]+)\?/.exec(e);
  if (!n)
    return [e];
  let t = e.slice(0, n.index), s = e.slice(n.index + n[0].length);
  const r = [t, t += n[1]];
  for (; n = /^(\/\:[^\/]+)\?/.exec(s); )
    r.push(t += n[1]), s = s.slice(n[0].length);
  return A(s).reduce((o, a) => [...o, ...r.map((i) => i + a)], []);
}
function ue(e) {
  const n = e.path.split("/").filter(Boolean), t = [], s = [];
  let r = 0, o = false;
  for (const [a, i] of n.entries())
    if (i[0] === ":") {
      const c = i.slice(1);
      r += 3, t.push({ type: ":", name: c, index: a }), s.push(null);
    } else
      i[0] === "*" ? (r -= 1, t.push({ type: "*", name: i.slice(1), index: a }), o = true) : (r += 4, i.match(/^\(.+\)$/) || s.push(i));
  return { ...e, score: r, params: t, matchSegments: s, wildcard: o };
}
function pe() {
  function e(t) {
    return { ...t, ...t.$$route ? t.$$route.require().route : void 0, metadata: { ...t.$$route ? t.$$route.require().route.metadata : {}, filesystem: true }, component: re(t.$component, globalThis.MANIFEST.client, globalThis.MANIFEST.ssr), children: t.children ? t.children.map(e) : void 0 };
  }
  return ae.map(e);
}
function de(e) {
  const n = getCookie(e, "flash");
  if (!n)
    return;
  let t = JSON.parse(n);
  if (!t || !t.result)
    return [];
  const s = [...t.input.slice(0, -1), new Map(t.input[t.input.length - 1])];
  return setCookie(e, "flash", "", { maxAge: 0 }), { url: t.url, result: t.error ? new Error(t.result) : t.result, input: s };
}
async function me(e) {
  const n = globalThis.MANIFEST.client;
  return globalThis.MANIFEST.ssr, setResponseHeader(e, "Content-Type", "text/html"), Object.assign(e, { manifest: await n.json(), assets: [...await n.inputs[n.handler].assets()], initialSubmission: de(e), routes: pe(), components: { status: (s) => (setResponseStatus(e, s.code, s.text), () => setResponseStatus(e, 200)), header: (s) => (s.append ? appendResponseHeader(e, s.name, s.value) : setResponseHeader(e, s.name, s.value), () => {
    const r = getResponseHeader(e, s.name);
    if (r && typeof r == "string") {
      const o = r.split(", "), a = o.indexOf(s.value);
      a !== -1 && o.splice(a, 1), o.length ? setResponseHeader(e, s.name, o.join(", ")) : removeResponseHeader(e, s.name);
    }
  }) }, $islands: /* @__PURE__ */ new Set() });
}
function fe(e, n = {}) {
  return eventHandler({ onRequest: n.onRequest, onBeforeResponse: n.onBeforeResponse, handler: (t) => {
    const s = Q(t);
    return provideRequestEvent(s, async () => {
      const r = ie(new URL(s.request.url).pathname, s.request.method);
      if (r) {
        const p = (await r.handler.import())[s.request.method];
        return s.params = r.params, sharedConfig.context = { event: s }, await p(s);
      }
      const o = await me(s);
      let a = { ...n };
      if (a.onCompleteAll) {
        const u = a.onCompleteAll;
        a.onCompleteAll = (p) => {
          $(o)(p), u(p);
        };
      } else
        a.onCompleteAll = $(o);
      if (a.onCompleteShell) {
        const u = a.onCompleteShell;
        a.onCompleteShell = (p) => {
          S(o, t)(), u(p);
        };
      } else
        a.onCompleteShell = S(o, t);
      const i = renderToStream(() => (sharedConfig.context.event = o, e(o)), a);
      if (o.response && o.response.headers.get("Location"))
        return sendRedirect(s, o.response.headers.get("Location"));
      const { writable: c, readable: l } = new TransformStream();
      return i.pipeTo(c), l;
    });
  } });
}
function S(e, n) {
  return () => {
    e.response && e.response.headers.get("Location") && (setResponseStatus(n, 302), setHeader(n, "Location", e.response.headers.get("Location")));
  };
}
function $(e) {
  return ({ write: n }) => {
    const t = e.response && e.response.headers.get("Location");
    t && n(`<script>window.location="${t}"<\/script>`);
  };
}
function he(e, n) {
  return fe(e, { ...n, createPageEvent: ge });
}
async function ge(e) {
  const n = globalThis.MANIFEST.client;
  return Object.assign(e, { manifest: await n.json(), assets: [...await n.inputs[n.handler].assets()], routes: [], $islands: /* @__PURE__ */ new Set() });
}
var ye = ['<head><script defer data-domain="ancgames.com" src="https://plausible.io/js/script.js"><\/script><meta charset="utf-8"><link rel="icon" href="/favicon.ico">', "</head>"], be = ["<html", ' lang="en">', '<body><div id="app">', "</div><!--$-->", "<!--/--></body></html>"];
const ve = he(() => createComponent(ne, { document: ({ assets: e, children: n, scripts: t }) => ssr(be, ssrHydrationKey(), createComponent(NoHydration, { get children() {
  return ssr(ye, escape(e));
} }), escape(n), escape(t)) }));

const handlers = [
  { route: '', handler: _f4b49z, lazy: false, middleware: true, method: undefined },
  { route: '/_server', handler: Z$1, lazy: false, middleware: true, method: undefined },
  { route: '/', handler: ve, lazy: false, middleware: true, method: undefined }
];

function createNitroApp() {
  const config = useRuntimeConfig();
  const hooks = createHooks();
  const captureError = (error, context = {}) => {
    const promise = hooks.callHookParallel("error", error, context).catch((_err) => {
      console.error("Error while capturing another error", _err);
    });
    if (context.event && isEvent(context.event)) {
      const errors = context.event.context.nitro?.errors;
      if (errors) {
        errors.push({ error, context });
      }
      if (context.event.waitUntil) {
        context.event.waitUntil(promise);
      }
    }
  };
  const h3App = createApp({
    debug: destr(false),
    onError: (error, event) => {
      captureError(error, { event, tags: ["request"] });
      return errorHandler(error, event);
    },
    onRequest: async (event) => {
      await nitroApp.hooks.callHook("request", event).catch((error) => {
        captureError(error, { event, tags: ["request"] });
      });
    },
    onBeforeResponse: async (event, response) => {
      await nitroApp.hooks.callHook("beforeResponse", event, response).catch((error) => {
        captureError(error, { event, tags: ["request", "response"] });
      });
    },
    onAfterResponse: async (event, response) => {
      await nitroApp.hooks.callHook("afterResponse", event, response).catch((error) => {
        captureError(error, { event, tags: ["request", "response"] });
      });
    }
  });
  const router = createRouter$1({
    preemptive: true
  });
  const localCall = createCall(toNodeListener(h3App));
  const _localFetch = createFetch(localCall, globalThis.fetch);
  const localFetch = (input, init) => _localFetch(input, init).then(
    (response) => normalizeFetchResponse(response)
  );
  const $fetch = createFetch$1({
    fetch: localFetch,
    Headers: Headers$1,
    defaults: { baseURL: config.app.baseURL }
  });
  globalThis.$fetch = $fetch;
  h3App.use(createRouteRulesHandler({ localFetch }));
  h3App.use(
    eventHandler((event) => {
      event.context.nitro = event.context.nitro || { errors: [] };
      const envContext = event.node.req?.__unenv__;
      if (envContext) {
        Object.assign(event.context, envContext);
      }
      event.fetch = (req, init) => fetchWithEvent(event, req, init, { fetch: localFetch });
      event.$fetch = (req, init) => fetchWithEvent(event, req, init, {
        fetch: $fetch
      });
      event.waitUntil = (promise) => {
        if (!event.context.nitro._waitUntilPromises) {
          event.context.nitro._waitUntilPromises = [];
        }
        event.context.nitro._waitUntilPromises.push(promise);
        if (envContext?.waitUntil) {
          envContext.waitUntil(promise);
        }
      };
      event.captureError = (error, context) => {
        captureError(error, { event, ...context });
      };
    })
  );
  for (const h of handlers) {
    let handler = h.lazy ? lazyEventHandler(h.handler) : h.handler;
    if (h.middleware || !h.route) {
      const middlewareBase = (config.app.baseURL + (h.route || "/")).replace(
        /\/+/g,
        "/"
      );
      h3App.use(middlewareBase, handler);
    } else {
      const routeRules = getRouteRulesForPath(
        h.route.replace(/:\w+|\*\*/g, "_")
      );
      if (routeRules.cache) {
        handler = cachedEventHandler(handler, {
          group: "nitro/routes",
          ...routeRules.cache
        });
      }
      router.use(h.route, handler, h.method);
    }
  }
  h3App.use(config.app.baseURL, router.handler);
  const app = {
    hooks,
    h3App,
    router,
    localCall,
    localFetch,
    captureError
  };
  for (const plugin of plugins) {
    try {
      plugin(app);
    } catch (err) {
      captureError(err, { tags: ["plugin"] });
      throw err;
    }
  }
  return app;
}
const nitroApp = createNitroApp();
const useNitroApp = () => nitroApp;

const localFetch = nitroApp.localFetch;
trapUnhandledNodeErrors();

export { localFetch };
//# sourceMappingURL=index.mjs.map
