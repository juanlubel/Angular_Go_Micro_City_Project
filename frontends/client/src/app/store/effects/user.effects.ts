import {Injectable} from '@angular/core';
import {Actions, Effect, ofType} from '@ngrx/effects';
import {select, Store} from '@ngrx/store';
import {of} from 'rxjs';
import {map, switchMap, withLatestFrom} from 'rxjs/operators';

import {IAppState} from '../state/app.state';
import {
  EUserActions,
  GetToken,
  GetTokenSuccess,
  GetUser,
  GetUsers,
  GetUsersSuccess,
  GetUserSuccess,
  UpdateUser
} from '../actions/user.actions';

import {UsersService} from '../../core/services';
import {IUserHttp, TokenHttp} from '../../models';
import {selectUserList} from '../selectors/user.selector';
import {TokenService} from '../../core/services/token.service';


@Injectable()
export class UserEffects {
  @Effect()
  getUser$ = this._actions$.pipe(
    ofType<GetUser>(EUserActions.GetUser),
    map(action => action.payload),
    withLatestFrom(this._store.pipe(select(selectUserList))),
    switchMap(([slug, users]) => {
      const selectedUser = users.filter(user => user.slug === slug)[0];
      console.log(selectedUser);
      return of(new GetUserSuccess(selectedUser));
    })
  );

  @Effect()
  getUsers$ = this._actions$.pipe(
    ofType<GetUsers>(EUserActions.GetUsers),
    switchMap(() => this._userService.getUsers()),
    switchMap((userHttp: IUserHttp) => {
      console.log(userHttp);
      return of(new GetUsersSuccess(userHttp.users));
    })
  );

  @Effect()
  updateUser$ = this._actions$.pipe(
    ofType<UpdateUser>(EUserActions.UpdateUser),
    map((action: any) => action.payload),
    switchMap((user) => this._userService.updateUser(user)),
    switchMap(() => this._userService.getUsers()),
    switchMap((userHttp: IUserHttp) => {
      console.log(userHttp);
      return of(new GetUsersSuccess(userHttp.users));
    })
  );

  @Effect()
  getToken$ = this._actions$.pipe(
    ofType<GetToken>(EUserActions.GetToken),
    map((action: any) => action.payload),
    switchMap( (payload) => this._tokenService.getToken(payload)),
    switchMap( (token: TokenHttp) => of(new GetTokenSuccess(token)))
  )

  constructor(
    private _userService: UsersService,
    private _tokenService: TokenService,
    private _actions$: Actions,
    private _store: Store<IAppState>
  ) {}
}
