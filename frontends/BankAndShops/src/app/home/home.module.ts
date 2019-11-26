import { NgModule } from "@angular/core";


import {HomeRoutingModule  } from "./home-routing.module";
import { HomeComponent } from "./home.component";
/* import { BanksComponent } from '../bank/banks.component'; */

@NgModule({
  imports: [ HomeRoutingModule],
  declarations: [HomeComponent,/* BanksComponent */]
})
export class HomeModule {}