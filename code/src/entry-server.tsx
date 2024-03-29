import { StartServer } from "@solidjs/start/server";
import { createHandler } from "@solidjs/start/entry";

export default createHandler(() => (
  <StartServer
    document={({ assets, children, scripts }) => (
      <html lang="en">
        <head>
          <script
            defer
            data-domain="ancgames.com"
            src="https://plausible.io/js/script.js"
          ></script>

          <meta charset="utf-8" />
          <link rel="icon" href="/favicon.ico" />
          {assets}
        </head>
        <body>
          <div id="app">{children}</div>
          {scripts}
        </body>
      </html>
    )}
  />
));
