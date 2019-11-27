import { NgModule } from "@angular/core";


import {HomeRoutingModule  } from "./home-routing.module";
import { HomeComponent } from "./home.component";
import { BankModule } from '../bank';


@NgModule({
  imports: [ HomeRoutingModule,BankModule],
  declarations: [HomeComponent]
})
export class HomeModule {}