import {RouterModule, Routes} from "@angular/router";
import {HomeComponent} from '../home/home.component';
import {NgModule} from "@angular/core";

const routes: Routes = [
  {path: '', component: HomeComponent},

];

@NgModule({
  exports: [RouterModule],
  imports: [RouterModule.forChild(routes)]
})
export class HomeRoutingModule {
}
