import { createEffect } from "solid-js";
import { SetStoreFunction, Store, createStore } from "solid-js/store";

export function createLocalStore<T>(
  initState: T,
  key: string
): [Store<T>, SetStoreFunction<T>] {
  const [state, setState] = createStore(initState as any);
  if (localStorage.getItem(key)) {
    try {
      setState(JSON.parse(localStorage.getItem(key) as string));
    } catch (error) {
      setState(() => initState);
    }
  }

  createEffect(() => {
    localStorage?.setItem(key, JSON.stringify(state));
  });

  return [state, setState];
}
