import React, { useState } from "react";

interface Props {
  children: React.ReactNode;
}

export interface AppContract {
    loggedIn: boolean,
    username?: string,
    apiToken?: string
}
    
export interface AppState {
    app?: AppContract
    updateState: (newState: Partial<AppState>) => void
}

const defaultState: AppState = {
    app: {loggedIn: false},
    updateState: (newState?: Partial<AppState>) => {},
}

export const AppContext = React.createContext<AppState>(defaultState)

export const AppContextProvider: React.FunctionComponent<Props> = (
  props: Props
): JSX.Element => {
  
  const [state, setState] = useState({});

  const updateState = (newState: Partial<AppState>) => {
    setState({ ...state, ...newState });
  }

  return (
    <AppContext.Provider value={{ ...state, updateState }}>
      {props.children}
    </AppContext.Provider>
  )
}