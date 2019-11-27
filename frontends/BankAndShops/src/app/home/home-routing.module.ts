import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";

import {  HomeComponent} from './home.component'
import { BankResolver } from '../bank';


const routes: Routes = [
    { path: "", component: HomeComponent,  /* resolve: {
        banks: BankResolver
      } */},

];

@NgModule({
    exports: [RouterModule],
    imports: [RouterModule.forChild(routes)]
})
export class HomeRoutingModule { }