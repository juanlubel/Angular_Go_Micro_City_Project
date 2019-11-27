import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";

import { BanksComponent } from './banks.component'
import { BankResolver } from './bank.resolver';


const routes: Routes = [
    {
        path: '',
        component: BanksComponent,
        resolve: {
          banks: BankResolver
        }
      }

];

@NgModule({
    exports: [RouterModule],
    imports: [RouterModule.forChild(routes)]
})
export class BankRoutingModule { }