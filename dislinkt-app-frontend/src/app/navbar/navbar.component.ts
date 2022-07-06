import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { environment } from '../../environments/environment';
import Swal from 'sweetalert2'
import axios from 'axios';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  constructor(private router: Router) { }

  notificationsAmount = 0;
  

  ngOnInit(): void {
    this.getNotifications()
  }

  getNotifications() {
    const id = localStorage.getItem('id');
    axios.get(environment.api + '/notifications/getNotifications/' + id)
      .then(response => {
        this.notificationsAmount = 0
        for(const p of (response.data as any)){
          if(!p.seen)
            this.notificationsAmount += 1
        }
      })
      .catch(e => {
        this.notificationsAmount = 0
        Swal.fire({
          icon: 'error',
          title: 'Something went wrong. Please, try again.',
          showConfirmButton: false,
          timer: 2000
        })
      })
  }

  logOut () {
    localStorage.removeItem('id')
    localStorage.removeItem('role')
    localStorage.removeItem('username')
    localStorage.removeItem('notificationsOn')
  }
}
