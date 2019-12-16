import {Component, Input, OnInit} from '@angular/core';
import {User} from '../../models';
import {UpdateUser} from '../../store/actions/user.actions';
import {Store} from '@ngrx/store';
import {IAppState} from '../../store/state/app.state';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css']
})
export class UserDetailsComponent implements OnInit {

  // tslint:disable-next-line:no-input-rename
  @Input('user') user: User;
  constructor(private store: Store<IAppState>) { }

  ngOnInit() {
  }

  updateUser(a, b) {
    console.log('update', this.user, a , b);
    const updateUser: User = this.user;
    updateUser.name = a;
    updateUser.email = b;
    this.store.dispatch(new UpdateUser(updateUser));
  }
}
