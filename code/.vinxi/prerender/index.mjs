globalThis._importMeta_={url:import.meta.url,env:process.env};import 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/node-fetch-native/dist/polyfill.cjs';
import { defineEventHandler, handleCacheHeaders, splitCookiesString, isEvent, createEvent, getRequestHeader, eventHandler, setHeaders, sendRedirect, proxyRequest, setResponseStatus, setResponseHeader, send, removeResponseHeader, createError, getResponseHeader, getHeader, getRequestURL, readFormData, readBody, setHeader, toWebRequest, getRequestIP, appendResponseHeader, getCookie, setCookie, createApp, createRouter as createRouter$1, toNodeListener, fetchWithEvent, lazyEventHandler } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/h3/dist/index.mjs';
import { createFetch as createFetch$1, Headers as Headers$1 } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/ofetch/dist/node.mjs';
import destr from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/destr/dist/index.mjs';
import { createCall, createFetch } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/unenv/runtime/fetch/index.mjs';
import { createHooks } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/hookable/dist/index.mjs';
import { snakeCase } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/scule/dist/index.mjs';
import { klona } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/klona/dist/index.mjs';
import defu, { defuFn } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/defu/dist/defu.mjs';
import { hash } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/ohash/dist/index.mjs';
import { parseURL, withoutBase, joinURL, getQuery, withQuery, decodePath, withLeadingSlash, withoutTrailingSlash } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/ufo/dist/index.mjs';
import { createStorage, prefixStorage } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/unstorage/dist/index.mjs';
import unstorage_47drivers_47fs from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/unstorage/drivers/fs.mjs';
import unstorage_47drivers_47fs_45lite from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/unstorage/drivers/fs-lite.mjs';
import { toRouteMatcher, createRouter } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/radix3/dist/index.mjs';
import _nkgJJC9dd4 from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/vinxi/lib/app-fetch.js';
import _qyzG6nrnQt from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/vinxi/lib/app-manifest.js';
import { promises } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { dirname, resolve } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/pathe/dist/index.mjs';
import { fromJSON, crossSerializeStream } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/seroval/dist/esm/production/index.mjs';
import { CustomEventPlugin, DOMExceptionPlugin, EventPlugin, FormDataPlugin, HeadersPlugin, ReadableStreamPlugin, RequestPlugin, ResponsePlugin, URLSearchParamsPlugin, URLPlugin } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/seroval-plugins/dist/esm/production/web.mjs';
import { provideRequestEvent } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/solid-js/web/dist/storage.js';
import { ssr, renderToStream, createComponent, ssrHydrationKey, NoHydration, escape, getRequestEvent, ssrAttribute, ssrElement, mergeProps } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/solid-js/web/dist/server.js';
import { lazy, createComponent as createComponent$1 } from 'file:///Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/node_modules/solid-js/dist/server.js';

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

const serverAssets = [{"baseName":"server","dir":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/assets"}];

const assets$1 = createStorage();

for (const asset of serverAssets) {
  assets$1.mount(asset.baseName, unstorage_47drivers_47fs({ base: asset.dir }));
}

const storage = createStorage({});

storage.mount('/assets', assets$1);

storage.mount('data', unstorage_47drivers_47fs_45lite({"driver":"fsLite","base":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/.data/kv"}));
storage.mount('root', unstorage_47drivers_47fs({"driver":"fs","readOnly":true,"base":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code","ignore":["**/node_modules/**","**/.git/**"]}));
storage.mount('src', unstorage_47drivers_47fs({"driver":"fs","readOnly":true,"base":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code","ignore":["**/node_modules/**","**/.git/**"]}));
storage.mount('build', unstorage_47drivers_47fs({"driver":"fs","readOnly":false,"base":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/.vinxi","ignore":["**/node_modules/**","**/.git/**"]}));
storage.mount('cache', unstorage_47drivers_47fs({"driver":"fs","readOnly":false,"base":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/.vinxi/cache","ignore":["**/node_modules/**","**/.git/**"]}));

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

const appConfig = {"name":"vinxi","routers":[{"name":"public","mode":"static","dir":"./public","base":"/","root":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code","order":0,"outDir":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/.vinxi/build/public"},{"name":"ssr","mode":"handler","handler":"src/entry-server.tsx","extensions":["js","jsx","ts","tsx"],"target":"server","root":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code","base":"/","outDir":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/.vinxi/build/ssr","order":1},{"name":"client","mode":"build","handler":"src/entry-client.tsx","extensions":["js","jsx","ts","tsx"],"target":"browser","base":"/_build","root":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code","outDir":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/.vinxi/build/client","order":2},{"name":"server-fns","mode":"handler","base":"/_server","handler":"node_modules/@solidjs/start/config/server-handler.js","target":"server","root":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code","outDir":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/.vinxi/build/server-fns","order":3}],"server":{"compressPublicAssets":{"brotli":true},"base":"./","prerender":{}},"root":"/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code"};
				const buildManifest = {"ssr":{"src/routes/index.tsx?pick=default&pick=$css":{"file":"index.js","isDynamicEntry":true,"isEntry":true,"src":"src/routes/index.tsx?pick=default&pick=$css"},"virtual:#vinxi/handler/ssr":{"dynamicImports":["src/routes/index.tsx?pick=default&pick=$css","src/routes/index.tsx?pick=default&pick=$css"],"file":"ssr.js","isEntry":true,"src":"virtual:#vinxi/handler/ssr"}},"client":{"\u0000virtual:#vinxi/handler/client.css":{"file":"assets/client-fab64f01.css","src":"\u0000virtual:#vinxi/handler/client.css"},"_index-76e275ee.js":{"file":"assets/index-76e275ee.js"},"src/routes/index.tsx?pick=default&pick=$css":{"file":"assets/index-23b670f1.js","imports":["_index-76e275ee.js"],"isDynamicEntry":true,"isEntry":true,"src":"src/routes/index.tsx?pick=default&pick=$css"},"virtual:#vinxi/handler/client":{"css":["assets/client-fab64f01.css"],"dynamicImports":["src/routes/index.tsx?pick=default&pick=$css"],"file":"assets/client-8a624d71.js","imports":["_index-76e275ee.js"],"isEntry":true,"src":"virtual:#vinxi/handler/client"}},"server-fns":{"virtual:#vinxi/handler/server-fns":{"file":"entry.js","isEntry":true,"src":"virtual:#vinxi/handler/server-fns"}}};

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
_nkgJJC9dd4,
_qyzG6nrnQt,
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
  "/favicon.ico": {
    "type": "image/vnd.microsoft.icon",
    "etag": "\"2939-3fRK9BhMjHtxjc3rXBuaoU1HteA\"",
    "mtime": "2024-01-07T20:02:44.254Z",
    "size": 10553,
    "path": "../../.output/public/favicon.ico"
  },
  "/_build/manifest.json": {
    "type": "application/json",
    "etag": "\"328-Yro/xh8dVTCYcAP1sYd3BVEMrQM\"",
    "mtime": "2024-01-07T20:02:44.255Z",
    "size": 808,
    "path": "../../.output/public/_build/manifest.json"
  },
  "/_build/server-functions-manifest.json": {
    "type": "application/json",
    "etag": "\"19-U+evudgPW1yE9kGumdxd/vtvk2s\"",
    "mtime": "2024-01-07T20:02:44.255Z",
    "size": 25,
    "path": "../../.output/public/_build/server-functions-manifest.json"
  },
  "/_build/assets/client-8a624d71.js": {
    "type": "application/javascript",
    "etag": "\"4359-ChYVwElnQtzls/Ak8qLQepmNjNs\"",
    "mtime": "2024-01-07T20:02:44.255Z",
    "size": 17241,
    "path": "../../.output/public/_build/assets/client-8a624d71.js"
  },
  "/_build/assets/client-8a624d71.js.br": {
    "type": "application/javascript",
    "encoding": "br",
    "etag": "\"1943-OsWLeqm304H7tXLZY4iqLyLxltg\"",
    "mtime": "2024-01-07T20:02:44.274Z",
    "size": 6467,
    "path": "../../.output/public/_build/assets/client-8a624d71.js.br"
  },
  "/_build/assets/client-8a624d71.js.gz": {
    "type": "application/javascript",
    "encoding": "gzip",
    "etag": "\"1c25-0K7MbavO+fn/9egSiFiRUidp1S0\"",
    "mtime": "2024-01-07T20:02:44.274Z",
    "size": 7205,
    "path": "../../.output/public/_build/assets/client-8a624d71.js.gz"
  },
  "/_build/assets/client-fab64f01.css": {
    "type": "text/css; charset=utf-8",
    "etag": "\"2568-1w5LcC9g82fbeg9wyETLjCV1eHI\"",
    "mtime": "2024-01-07T20:02:44.255Z",
    "size": 9576,
    "path": "../../.output/public/_build/assets/client-fab64f01.css"
  },
  "/_build/assets/client-fab64f01.css.br": {
    "type": "text/css; charset=utf-8",
    "encoding": "br",
    "etag": "\"7f5-WJ5mQSsMkyqsnrCwWTt6shSCRGI\"",
    "mtime": "2024-01-07T20:02:44.274Z",
    "size": 2037,
    "path": "../../.output/public/_build/assets/client-fab64f01.css.br"
  },
  "/_build/assets/client-fab64f01.css.gz": {
    "type": "text/css; charset=utf-8",
    "encoding": "gzip",
    "etag": "\"96b-j5z1zPfe1Cu68WYHFlOldF8Nqj8\"",
    "mtime": "2024-01-07T20:02:44.274Z",
    "size": 2411,
    "path": "../../.output/public/_build/assets/client-fab64f01.css.gz"
  },
  "/_build/assets/index-23b670f1.js": {
    "type": "application/javascript",
    "etag": "\"37c68-/Pk74CO3H0b3S19eMTtc2/TB/cg\"",
    "mtime": "2024-01-07T20:02:44.256Z",
    "size": 228456,
    "path": "../../.output/public/_build/assets/index-23b670f1.js"
  },
  "/_build/assets/index-23b670f1.js.br": {
    "type": "application/javascript",
    "encoding": "br",
    "etag": "\"13bc6-f+ef8hYzU7zmd6Ubxp1a+wf4is4\"",
    "mtime": "2024-01-07T20:02:44.529Z",
    "size": 80838,
    "path": "../../.output/public/_build/assets/index-23b670f1.js.br"
  },
  "/_build/assets/index-23b670f1.js.gz": {
    "type": "application/javascript",
    "encoding": "gzip",
    "etag": "\"162ec-5uBaLY97+Av2oEdhygwkI4JwFo8\"",
    "mtime": "2024-01-07T20:02:44.302Z",
    "size": 90860,
    "path": "../../.output/public/_build/assets/index-23b670f1.js.gz"
  },
  "/_build/assets/index-76e275ee.js": {
    "type": "application/javascript",
    "etag": "\"5e73-tjmCGSOIyREcb4n2vXadQ1fv/uQ\"",
    "mtime": "2024-01-07T20:02:44.255Z",
    "size": 24179,
    "path": "../../.output/public/_build/assets/index-76e275ee.js"
  },
  "/_build/assets/index-76e275ee.js.br": {
    "type": "application/javascript",
    "encoding": "br",
    "etag": "\"21de-t1awElCu/0M3CXXtd9yVO3Qoowc\"",
    "mtime": "2024-01-07T20:02:44.290Z",
    "size": 8670,
    "path": "../../.output/public/_build/assets/index-76e275ee.js.br"
  },
  "/_build/assets/index-76e275ee.js.gz": {
    "type": "application/javascript",
    "encoding": "gzip",
    "etag": "\"24ff-BCj20GjONoJFAoQi7BTOjiYz+Mk\"",
    "mtime": "2024-01-07T20:02:44.274Z",
    "size": 9471,
    "path": "../../.output/public/_build/assets/index-76e275ee.js.gz"
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
const p = "Invariant Violation", { setPrototypeOf: z = function(e, t) {
  return e.__proto__ = t, e;
} } = Object;
class f extends Error {
  constructor(t = p) {
    super(typeof t == "number" ? `${p}: ${t} (see https://github.com/apollographql/invariant-packages)` : t);
    __publicField(this, "framesToPop", 1);
    __publicField(this, "name", p);
    z(this, f.prototype);
  }
}
function h$1(e, t) {
  if (!e)
    throw new f(t);
}
const T$1 = Symbol("h3Event"), l = Symbol("fetchEvent"), J$1 = { get(e, t) {
  var _a;
  return t === l ? e : (_a = e[t]) != null ? _a : e[T$1][t];
} };
function M(e) {
  return new Proxy({ request: toWebRequest(e), clientAddress: getRequestIP(e), locals: {}, [T$1]: e }, J$1);
}
function W$1(e) {
  if (!e[l]) {
    const t = M(e);
    e[l] = t;
  }
  return e[l];
}
function d(e, t) {
  return new ReadableStream({ start(s) {
    crossSerializeStream(t, { scopeId: e, plugins: [CustomEventPlugin, DOMExceptionPlugin, EventPlugin, FormDataPlugin, HeadersPlugin, ReadableStreamPlugin, RequestPlugin, ResponsePlugin, URLSearchParamsPlugin, URLPlugin], onSerialize(n, a) {
      const i = a ? `($R["${e}"]=[],${n})` : n;
      s.enqueue(new TextEncoder().encode(`${i};
`));
    }, onDone() {
      s.close();
    }, onError(n) {
      s.error(n);
    } });
  } });
}
async function _(e) {
  h$1(e.method === "POST", `Invalid method ${e.method}. Expected POST.`);
  const t = getHeader(e, "x-server-id"), s = getHeader(e, "x-server-instance"), n = getRequestURL(e);
  let a, i;
  if (t)
    h$1(typeof t == "string", "Invalid server function"), [a, i] = t.split("#");
  else if (a = n.searchParams.get("id"), i = n.searchParams.get("name"), !a || !i)
    throw new Error("Invalid request");
  const q = (await globalThis.MANIFEST["server-fns"].chunks[a].import())[i];
  let o = [];
  if (!s) {
    const r = n.searchParams.get("args");
    r && JSON.parse(r).forEach((c) => o.push(c));
  }
  const m = getHeader(e, "content-type");
  m.startsWith("multipart/form-data") || m.startsWith("application/x-www-form-urlencoded") ? o.push(await readFormData(e)) : o = fromJSON(await readBody(e), { plugins: [CustomEventPlugin, DOMExceptionPlugin, EventPlugin, FormDataPlugin, HeadersPlugin, ReadableStreamPlugin, RequestPlugin, ResponsePlugin, URLSearchParamsPlugin, URLPlugin] });
  try {
    const r = await provideRequestEvent(W$1(e), () => q(...o));
    if (!s) {
      const c = r instanceof Error, O = new URL(getHeader(e, "referer"));
      return new Response(null, { status: 302, headers: { Location: O.toString(), ...r ? { "Set-Cookie": `flash=${JSON.stringify({ url: n.pathname + encodeURIComponent(n.search), result: c ? r.message : r, error: c, input: [...o.slice(0, -1), [...o[o.length - 1].entries()]] })}; Secure; HttpOnly;` } : {} } });
    }
    return setHeader(e, "content-type", "text/javascript"), d(s, r);
  } catch (r) {
    return r instanceof Response && r.status === 302 ? new Response(null, { status: s ? 204 : 302, headers: { Location: r.headers.get("Location") } }) : new Response(d(s, r), { status: 500, headers: { "Content-Type": "text/javascript" } });
  }
}
const G$1 = eventHandler(_);

const b = [{ type: "page", $component: { src: "src/routes/index.tsx?pick=default&pick=$css", build: () => import('./chunks/build/index.mjs'), import: () => import('./chunks/build/index.mjs') }, path: "/", filePath: "/Users/andrew/Documents/Dev/Projects/andrewnathanco/github/challenge/code/src/routes/index.tsx" }], J = K(b.filter((e) => e.type === "page")), U = Y(b.filter((e) => e.type === "api"));
function W(e, n) {
  const t = e.split("/").filter(Boolean);
  e:
    for (const s of U) {
      const r = s.matchSegments;
      if (t.length < r.length || !s.wildcard && t.length > r.length)
        continue;
      for (let i = 0; i < r.length; i++) {
        const l = r[i];
        if (l && t[i] !== l)
          continue e;
      }
      const a = s[`$${n}`];
      if (a === "skip" || a === void 0)
        return;
      const o = {};
      for (const { type: i, name: l, index: c } of s.params)
        i === ":" ? o[l] = t[c] : o[l] = t.slice(c).join("/");
      return { handler: a, params: o };
    }
}
function K(e) {
  function n(t, s, r, a) {
    const o = Object.values(t).find((i) => r.startsWith(i.id + "/"));
    return o ? (n(o.children || (o.children = []), s, r.slice(o.id.length)), t) : (t.push({ ...s, id: r, path: r.replace(/\/\([^)/]+\)/g, "") }), t);
  }
  return e.sort((t, s) => t.path.length - s.path.length).reduce((t, s) => n(t, s, s.path, s.path), []);
}
function Y(e) {
  return e.flatMap((n) => C(n.path).map((s) => ({ ...n, path: s }))).map(G).sort((n, t) => t.score - n.score);
}
function C(e) {
  let n = /(\/?\:[^\/]+)\?/.exec(e);
  if (!n)
    return [e];
  let t = e.slice(0, n.index), s = e.slice(n.index + n[0].length);
  const r = [t, t += n[1]];
  for (; n = /^(\/\:[^\/]+)\?/.exec(s); )
    r.push(t += n[1]), s = s.slice(n[0].length);
  return C(s).reduce((a, o) => [...a, ...r.map((i) => i + o)], []);
}
function G(e) {
  const n = e.path.split("/").filter(Boolean), t = [], s = [];
  let r = 0, a = false;
  for (const [o, i] of n.entries())
    if (i[0] === ":") {
      const l = i.slice(1);
      r += 3, t.push({ type: ":", name: l, index: o }), s.push(null);
    } else
      i[0] === "*" ? (r -= 1, t.push({ type: "*", name: i.slice(1), index: o }), a = true) : (r += 4, s.push(i));
  return { ...e, score: r, params: t, matchSegments: s, wildcard: a };
}
const T = Symbol("h3Event"), h = Symbol("fetchEvent"), Q = { get(e, n) {
  var _a;
  return n === h ? e : (_a = e[n]) != null ? _a : e[T][n];
} };
function V(e) {
  return new Proxy({ request: toWebRequest(e), clientAddress: getRequestIP(e), locals: {}, [T]: e }, Q);
}
function X(e) {
  if (!e[h]) {
    const n = V(e);
    e[h] = n;
  }
  return e[h];
}
const Z = " ", ee = { style: (e) => ssrElement("style", e.attrs, () => escape(e.children), true), link: (e) => ssrElement("link", e.attrs, void 0, true), script: (e) => e.attrs.src ? ssrElement("script", mergeProps(() => e.attrs, { get id() {
  return e.key;
} }), () => ssr(Z), true) : null };
function E(e) {
  let { tag: n, attrs: { key: t, ...s } = { key: void 0 }, children: r } = e;
  return ee[n]({ attrs: s, key: t, children: r });
}
function te(e, n, t, s = "default") {
  return lazy(async () => {
    var _a;
    {
      const a = (await e.import())[s], i = (await ((_a = n.inputs) == null ? void 0 : _a[e.src].assets())).filter((c) => c.tag === "style" || c.attrs.rel === "stylesheet");
      return { default: (c) => [...i.map((p) => E(p)), createComponent$1(a, c)] };
    }
  });
}
function ne() {
  function e(t) {
    return { ...t, ...t.$$route ? t.$$route.require().route : void 0, metadata: { ...t.$$route ? t.$$route.require().route.metadata : {}, filesystem: true }, component: te(t.$component, globalThis.MANIFEST.client, globalThis.MANIFEST.ssr), children: t.children ? t.children.map(e) : void 0 };
  }
  return J.map(e);
}
function se(e) {
  const n = getCookie(e, "flash");
  if (!n)
    return;
  let t = JSON.parse(n);
  if (!t || !t.result)
    return [];
  const s = [...t.input.slice(0, -1), new Map(t.input[t.input.length - 1])];
  return setCookie(e, "flash", "", { maxAge: 0 }), { url: t.url, result: t.error ? new Error(t.result) : t.result, input: s };
}
async function re(e) {
  const n = globalThis.MANIFEST.client;
  return globalThis.MANIFEST.ssr, setResponseHeader(e, "Content-Type", "text/html"), Object.assign(e, { manifest: await n.json(), assets: [...await n.inputs[n.handler].assets()], initialSubmission: se(e), routes: ne(), components: { status: (s) => (setResponseStatus(e, s.code, s.text), () => setResponseStatus(e, 200)), header: (s) => (s.append ? appendResponseHeader(e, s.name, s.value) : setResponseHeader(e, s.name, s.value), () => {
    const r = getResponseHeader(e, s.name);
    if (r && typeof r == "string") {
      const a = r.split(", "), o = a.indexOf(s.value);
      o !== -1 && a.splice(o, 1), a.length ? setResponseHeader(e, s.name, a.join(", ")) : removeResponseHeader(e, s.name);
    }
  }) }, $islands: /* @__PURE__ */ new Set() });
}
function oe(e, n = {}) {
  return eventHandler({ onRequest: n.onRequest, onBeforeResponse: n.onBeforeResponse, handler: (t) => {
    const s = X(t);
    return provideRequestEvent(s, async () => {
      const r = W(new URL(s.request.url).pathname, s.request.method);
      if (r) {
        const m = (await r.handler.import())[s.request.method];
        return s.params = r.params, await m(s);
      }
      const a = await re(s);
      let o = { ...n };
      if (o.onCompleteAll) {
        const p = o.onCompleteAll;
        o.onCompleteAll = (m) => {
          v(a)(m), p(m);
        };
      } else
        o.onCompleteAll = v(a);
      if (o.onCompleteShell) {
        const p = o.onCompleteShell;
        o.onCompleteShell = (m) => {
          $(a, t)(), p(m);
        };
      } else
        o.onCompleteShell = $(a, t);
      const i = renderToStream(() => e(a), o);
      if (a.response && a.response.headers.get("Location"))
        return sendRedirect(s, a.response.headers.get("Location"));
      const { writable: l, readable: c } = new TransformStream();
      return i.pipeTo(l), c;
    });
  } });
}
function $(e, n) {
  return () => {
    e.response && e.response.headers.get("Location") && (setResponseStatus(n, 302), setHeader(n, "Location", e.response.headers.get("Location")));
  };
}
function v(e) {
  return ({ write: n }) => {
    const t = e.response && e.response.headers.get("Location");
    t && n(`<script>window.location="${t}"<\/script>`);
  };
}
const ae = ["<script", ">$R = [];<\/script>"], ie = ["<script", ">", "<\/script>"], le = ["<script", ' type="module"', "><\/script>"], ce = ssr("<!DOCTYPE html>");
function ue(e) {
  const n = getRequestEvent();
  return createComponent(NoHydration, { get children() {
    return [ce, createComponent(e.document, { get assets() {
      return [ssr(ae, ssrHydrationKey()), n.assets.map((t) => E(t))];
    }, get scripts() {
      return [ssr(ie, ssrHydrationKey(), `window.manifest = ${JSON.stringify(n.manifest)}`), ssr(le, ssrHydrationKey(), ssrAttribute("src", escape(globalThis.MANIFEST.client.inputs[globalThis.MANIFEST.client.handler].output.path, true), false))];
    } })];
  } });
}
const pe = ['<head><meta charset="utf-8"><link rel="icon" href="/favicon.ico">', "</head>"], me = ["<html", ' lang="en">', '<body><div id="app">', "</div><!--$-->", "<!--/--></body></html>"], ye = oe(() => createComponent(ue, { document: ({ assets: e, children: n, scripts: t }) => ssr(me, ssrHydrationKey(), createComponent(NoHydration, { get children() {
  return ssr(pe, escape(e));
} }), escape(n), escape(t)) }));

const handlers = [
  { route: '', handler: _f4b49z, lazy: false, middleware: true, method: undefined },
  { route: '/_server', handler: G$1, lazy: false, middleware: true, method: undefined },
  { route: '/', handler: ye, lazy: false, middleware: true, method: undefined }
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
