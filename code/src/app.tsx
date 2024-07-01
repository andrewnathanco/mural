// @refresh reload
import { Router } from "@solidjs/router";
import { Suspense } from "solid-js";
import { FileRoutes } from "@solidjs/start/router";
import { Meta, MetaProvider, Title } from "@solidjs/meta";
import "./app.css";

export default function App() {
  return (
    <MetaProvider>
      <Title>Mural</Title>
      <Meta name="description" content="The daily mural" />
      <Meta
        name="keywords"
        content="ANCGames, daily puzzles, puzzle games, brain games, online puzzles, daily challenge, wordle, nytgames, the crossword, nyt, anc"
      />
      <Meta name="author" content="Andrew Cohen" />
      <Router
        base={import.meta.env.SERVER_BASE_URL}
        root={(props) => (
          <>
            <Suspense>{props.children}</Suspense>
          </>
        )}
      >
        <FileRoutes />
      </Router>
    </MetaProvider>
  );
}
