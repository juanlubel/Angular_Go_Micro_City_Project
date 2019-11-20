import {Component, OnInit} from '@angular/core';
import {UsersService} from "../services/users.service";

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  title: string = "Hello Users"
  dataParent: string[] = ["uno", "dos", "tres"]
  hostedString: string

  constructor(private usersService: UsersService) {
  }

  ngOnInit() {
    console.log(this.usersService.get().subscribe((res: any) => {
      console.log(res)
    }))
  }

  eventListener(event) {
    console.log(event);
  }
}
