// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { CircleDB } from './circle-db';
import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { AnimateDB } from './animate-db'

@Injectable({
  providedIn: 'root'
})
export class CircleService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  CircleServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private circlesUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.circlesUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/circles';
  }

  /** GET circles from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CircleDB[]> {
    return this.getCircles(GONG__StackPath, frontRepo)
  }
  getCircles(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CircleDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<CircleDB[]>(this.circlesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<CircleDB[]>('getCircles', []))
      );
  }

  /** GET circle by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CircleDB> {
    return this.getCircle(id, GONG__StackPath, frontRepo)
  }
  getCircle(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CircleDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.circlesUrl}/${id}`;
    return this.http.get<CircleDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched circle id=${id}`)),
      catchError(this.handleError<CircleDB>(`getCircle id=${id}`))
    );
  }

  /** POST: add a new circle to the server */
  post(circledb: CircleDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CircleDB> {
    return this.postCircle(circledb, GONG__StackPath, frontRepo)
  }
  postCircle(circledb: CircleDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CircleDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    circledb.CirclePointersEncoding.Animations = []
    for (let _animate of circledb.Animations) {
      circledb.CirclePointersEncoding.Animations.push(_animate.ID)
    }
    circledb.Animations = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CircleDB>(this.circlesUrl, circledb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        circledb.Animations = new Array<AnimateDB>()
        for (let _id of circledb.CirclePointersEncoding.Animations) {
          let _animate = frontRepo.Animates.get(_id)
          if (_animate != undefined) {
            circledb.Animations.push(_animate!)
          }
        }
        // this.log(`posted circledb id=${circledb.ID}`)
      }),
      catchError(this.handleError<CircleDB>('postCircle'))
    );
  }

  /** DELETE: delete the circledb from the server */
  delete(circledb: CircleDB | number, GONG__StackPath: string): Observable<CircleDB> {
    return this.deleteCircle(circledb, GONG__StackPath)
  }
  deleteCircle(circledb: CircleDB | number, GONG__StackPath: string): Observable<CircleDB> {
    const id = typeof circledb === 'number' ? circledb : circledb.ID;
    const url = `${this.circlesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<CircleDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted circledb id=${id}`)),
      catchError(this.handleError<CircleDB>('deleteCircle'))
    );
  }

  /** PUT: update the circledb on the server */
  update(circledb: CircleDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CircleDB> {
    return this.updateCircle(circledb, GONG__StackPath, frontRepo)
  }
  updateCircle(circledb: CircleDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CircleDB> {
    const id = typeof circledb === 'number' ? circledb : circledb.ID;
    const url = `${this.circlesUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers
    circledb.CirclePointersEncoding.Animations = []
    for (let _animate of circledb.Animations) {
      circledb.CirclePointersEncoding.Animations.push(_animate.ID)
    }
    circledb.Animations = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<CircleDB>(url, circledb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        circledb.Animations = new Array<AnimateDB>()
        for (let _id of circledb.CirclePointersEncoding.Animations) {
          let _animate = frontRepo.Animates.get(_id)
          if (_animate != undefined) {
            circledb.Animations.push(_animate!)
          }
        }
        // this.log(`updated circledb id=${circledb.ID}`)
      }),
      catchError(this.handleError<CircleDB>('updateCircle'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in CircleService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("CircleService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}
