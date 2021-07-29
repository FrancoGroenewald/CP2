import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { LoginService } from '../services/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  public isLoading: boolean = false;
  constructor(private logService: LoginService) { }

  ngOnInit(): void {
  }

  public onLogin = (form: NgForm) => {
    this.logService.login(form.value)
    .subscribe((response) => {
      console.log(response)
    })
  }
}
