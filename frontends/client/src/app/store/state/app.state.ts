import {RouterReducerState} from '@ngrx/router-store';
import {IUserState} from './user.state';


export interface IAppState {
  router?: RouterReducerState;
  // admins: AdminState;
  users: IUserState;
}

export const initialAppState: IAppState = {
  // admins: null,
  users: null,
};

export function getInitialState(): IAppState {
  return initialAppState;
}

