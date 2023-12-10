import { Accessor, createContext, createSignal, useContext } from "solid-js";

type ShareDialog = [
  Accessor<boolean>,
  {
    close: () => void;
    open: () => void;
  }
];

const ShareDialogContext = createContext<ShareDialog>();

export function ShareDialogProvider(props: { children: any }) {
  const [dialog_status, set_dialog_status] = createSignal(false);
  const dialog: ShareDialog = [
    dialog_status,
    {
      close() {
        set_dialog_status(false);
        document.body.style.overflowY = "auto";
      },
      open() {
        set_dialog_status(true);
        document.body.style.position = "relative";
        document.body.style.overflowY = "hidden";
      },
    },
  ];

  return (
    <ShareDialogContext.Provider value={dialog}>
      {props.children}
    </ShareDialogContext.Provider>
  );
}

export function useShareDialog(): ShareDialog {
  return useContext(ShareDialogContext) as ShareDialog;
}
