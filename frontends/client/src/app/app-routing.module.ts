import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';


const routes: Routes = [
  // {
  //   path: '', pathMatch: 'full' , redirectTo: '/login'
  // },
  // {
  //   path: 'login',
  //   loadChildren: './login/login.module#LoginModule'
  // },
  // {
  //   path: 'users',
  //   loadChildren: './users/users.module#UsersModule'
  // },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})


export class AppRoutingModule {
}
