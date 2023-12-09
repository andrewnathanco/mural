import { Accessor, createContext, createSignal, useContext } from "solid-js";

type InfoDialog = [
  Accessor<boolean>,
  {
    close: () => void;
    open: () => void;
  }
];

const InfoDialogContext = createContext<InfoDialog>();

export function InfoDialogProvider(props: { children: any }) {
  const [dialog_status, set_dialog_status] = createSignal(false);
  const dialog: InfoDialog = [
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
    <InfoDialogContext.Provider value={dialog}>
      {props.children}
    </InfoDialogContext.Provider>
  );
}

export function useInfoDialog(): InfoDialog {
  return useContext(InfoDialogContext) as InfoDialog;
}
