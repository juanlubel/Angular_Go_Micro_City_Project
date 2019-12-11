import {Component, OnInit} from '@angular/core';
import {UsersService} from '../core/services';
import {Store, select} from '@ngrx/store';
import {GetUsers} from '../store/actions/user.actions';
import {IAppState} from '../store/state/app.state';
import {selectUserList} from '../store/selectors/user.selector';
import {Observable} from 'rxjs';
import {User} from '../models';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  title = 'Hello Users';
  dataParent: string[] = ['uno', 'dos', 'tres'];
  hostedString: string;
  public users$ = new Observable<User[]>();

  constructor(private store: Store<IAppState>) {
    this.users$ = store.select(selectUserList);
    console.log(this.users$);
  }

  ngOnInit() {
    this.store.dispatch(new GetUsers());
  }

  eventListener(event) {
    console.log(event);
  }
}
