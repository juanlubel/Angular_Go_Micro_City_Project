import {Action} from '@ngrx/store';
import {TokenHttp, TokenType, User as IUser} from '../../models';


export enum EUserActions {
  GetUsers = '[User] Get Users',
  GetUsersSuccess = '[User] Get Users Success',
  GetUser = '[User] Get User',
  GetUserSuccess = '[User] Get User Success',
  UpdateUser = '[User] Update User',
  UpdateUserSuccess = ' [User] Update User Success',
  GetToken = ' [User] Get Token',
  GetTokenSuccess = ' [User] Get Token Success',
}

export class GetUsers implements Action {
  public readonly type = EUserActions.GetUsers;
}

export class GetUsersSuccess implements Action {
  public readonly type = EUserActions.GetUsersSuccess;

  constructor(public payload: IUser[]) {
  }
}

export class GetUser implements Action {
  public readonly type = EUserActions.GetUser;

  constructor(public payload: string) {
  }
}

export class GetUserSuccess implements Action {
  public readonly type = EUserActions.GetUserSuccess;

  constructor(public payload: IUser) {
  }
}

export class UpdateUser implements Action {
  public readonly type = EUserActions.UpdateUser;

  constructor(public payload: IUser) {
  }
}

export class UpdateUserSuccess implements Action {
  public readonly type = EUserActions.UpdateUserSuccess;

  constructor(public payload: IUser[]) {
  }
}

export class GetToken implements Action {
  public readonly type = EUserActions.GetToken;

  constructor(public payload: TokenType) {
  }
}

export class GetTokenSuccess implements Action {
  public readonly type = EUserActions.GetTokenSuccess;

  constructor(public payload: TokenHttp) {
  }
}

export type UserActions =
  GetUsers
  | GetUsersSuccess
  | GetUser
  | GetUserSuccess
  | UpdateUser
  | UpdateUserSuccess
  | GetToken
  | GetTokenSuccess;
