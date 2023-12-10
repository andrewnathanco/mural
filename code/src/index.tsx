import { render } from "solid-js/web";
/* @refresh reload */
import "./index.css";

import { Route, Router } from "@solidjs/router";
import Index from "./view";
import Share from "./view/share";

const root = document.getElementById("root");

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    "Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?"
  );
}

render(
  () => (
    <Router>
      <Route path="/" component={Index} />
      <Route path="/share" component={Share} />
    </Router>
  ),
  root!
);
