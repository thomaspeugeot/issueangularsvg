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

import { DisplayedColumnDB } from './displayedcolumn-db';
import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class DisplayedColumnService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  DisplayedColumnServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private displayedcolumnsUrl: string

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
    this.displayedcolumnsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/displayedcolumns';
  }

  /** GET displayedcolumns from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<DisplayedColumnDB[]> {
    return this.getDisplayedColumns(GONG__StackPath, frontRepo)
  }
  getDisplayedColumns(GONG__StackPath: string, frontRepo: FrontRepo): Observable<DisplayedColumnDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<DisplayedColumnDB[]>(this.displayedcolumnsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<DisplayedColumnDB[]>('getDisplayedColumns', []))
      );
  }

  /** GET displayedcolumn by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DisplayedColumnDB> {
    return this.getDisplayedColumn(id, GONG__StackPath, frontRepo)
  }
  getDisplayedColumn(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DisplayedColumnDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.displayedcolumnsUrl}/${id}`;
    return this.http.get<DisplayedColumnDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched displayedcolumn id=${id}`)),
      catchError(this.handleError<DisplayedColumnDB>(`getDisplayedColumn id=${id}`))
    );
  }

  /** POST: add a new displayedcolumn to the server */
  post(displayedcolumndb: DisplayedColumnDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DisplayedColumnDB> {
    return this.postDisplayedColumn(displayedcolumndb, GONG__StackPath, frontRepo)
  }
  postDisplayedColumn(displayedcolumndb: DisplayedColumnDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DisplayedColumnDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<DisplayedColumnDB>(this.displayedcolumnsUrl, displayedcolumndb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted displayedcolumndb id=${displayedcolumndb.ID}`)
      }),
      catchError(this.handleError<DisplayedColumnDB>('postDisplayedColumn'))
    );
  }

  /** DELETE: delete the displayedcolumndb from the server */
  delete(displayedcolumndb: DisplayedColumnDB | number, GONG__StackPath: string): Observable<DisplayedColumnDB> {
    return this.deleteDisplayedColumn(displayedcolumndb, GONG__StackPath)
  }
  deleteDisplayedColumn(displayedcolumndb: DisplayedColumnDB | number, GONG__StackPath: string): Observable<DisplayedColumnDB> {
    const id = typeof displayedcolumndb === 'number' ? displayedcolumndb : displayedcolumndb.ID;
    const url = `${this.displayedcolumnsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<DisplayedColumnDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted displayedcolumndb id=${id}`)),
      catchError(this.handleError<DisplayedColumnDB>('deleteDisplayedColumn'))
    );
  }

  /** PUT: update the displayedcolumndb on the server */
  update(displayedcolumndb: DisplayedColumnDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DisplayedColumnDB> {
    return this.updateDisplayedColumn(displayedcolumndb, GONG__StackPath, frontRepo)
  }
  updateDisplayedColumn(displayedcolumndb: DisplayedColumnDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DisplayedColumnDB> {
    const id = typeof displayedcolumndb === 'number' ? displayedcolumndb : displayedcolumndb.ID;
    const url = `${this.displayedcolumnsUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<DisplayedColumnDB>(url, displayedcolumndb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated displayedcolumndb id=${displayedcolumndb.ID}`)
      }),
      catchError(this.handleError<DisplayedColumnDB>('updateDisplayedColumn'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in DisplayedColumnService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("DisplayedColumnService" + error); // log to console instead

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
