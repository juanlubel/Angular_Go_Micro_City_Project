import {Component, OnInit} from '@angular/core';
import {Store} from '@ngrx/store';
import {GetUser, GetUsers, UpdateUser} from '../store/actions/user.actions';
import {IAppState} from '../store/state/app.state';
import {selectSelectedUser, selectUserList} from '../store/selectors/user.selector';
import {Observable} from 'rxjs';
import {User} from '../models';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  public users$ = new Observable<User[]>();
  public user$ = new Observable<User>();

  constructor(private store: Store<IAppState>) {
    this.users$ = store.select(selectUserList);
    this.user$ = this.store.select(selectSelectedUser);

  }

  ngOnInit() {
    this.store.dispatch(new GetUsers());
  }

  selectUser(slug: string) {
    console.log('click', slug);
    this.store.dispatch(new GetUser(slug));
    this.user$ = this.store.select(selectSelectedUser);
    console.log(this.user$);
  }

}
