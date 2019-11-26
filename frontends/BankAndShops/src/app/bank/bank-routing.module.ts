import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";

import { BanksComponent } from './banks.component'


const routes: Routes = [
    { path: "", component: BanksComponent},

];

@NgModule({
    exports: [RouterModule],
    imports: [RouterModule.forChild(routes)]
})
export class BankRoutingModule { }