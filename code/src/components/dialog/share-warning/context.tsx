import { Accessor, createContext, createSignal, useContext } from "solid-js";

type ShareWarningDialog = [
  Accessor<boolean>,
  {
    close: () => void;
    open: () => void;
  }
];

const ShareWarningDialogContext = createContext<ShareWarningDialog>();

export function ShareWarningDialogProvider(props: { children: any }) {
  const [dialog_status, set_dialog_status] = createSignal(false);
  const dialog: ShareWarningDialog = [
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
    <ShareWarningDialogContext.Provider value={dialog}>
      {props.children}
    </ShareWarningDialogContext.Provider>
  );
}

export function useShareWarningDialog(): ShareWarningDialog {
  return useContext(ShareWarningDialogContext) as ShareWarningDialog;
}
