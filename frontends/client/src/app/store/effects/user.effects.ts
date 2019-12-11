import {Injectable} from '@angular/core';
import {Actions, Effect, ofType} from '@ngrx/effects';
import {Store, select} from '@ngrx/store';
import {of} from 'rxjs';
import {map, switchMap, withLatestFrom} from 'rxjs/operators';

import {IAppState} from '../state/app.state';
import {
  EUserActions,
  GetUser,
  GetUsers,
  GetUsersSuccess,
  GetUserSuccess
} from '../actions/user.actions';

import {UsersService} from '../../core/services';
import {IUserHttp} from '../../models';
import {selectUserList} from '../selectors/user.selector';


@Injectable()
export class UserEffects {
  @Effect()
  getUser$ = this._actions$.pipe(
    ofType<GetUser>(EUserActions.GetUser),
    map(action => action.payload),
    withLatestFrom(this._store.pipe(select(selectUserList))),
    switchMap(([slug, users]) => {
      const selectedUser = users.filter(user => user.slug === slug)[0];
      return of(new GetUserSuccess(selectedUser));
    })
  );

  @Effect()
  getUsers$ = this._actions$.pipe(
    ofType<GetUsers>(EUserActions.GetUsers),
    switchMap(() => this._userService.getUsers()),
    switchMap((userHttp: IUserHttp) => of(new GetUsersSuccess(userHttp.users)))
  );

  constructor(
    private _userService: UsersService,
    private _actions$: Actions,
    private _store: Store<IAppState>
  ) {}
}
