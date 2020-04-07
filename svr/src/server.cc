#include "appweb.h"
#include <iostream>

static void doAction(HttpConn *conn)
{
    std::cout << "handling" << std::endl;
    HttpQueue *q;

    q = conn->writeq;
    /*
        Set the HTTP response status
     */
    httpSetStatus(conn, 200);

    /*
        Add desired headers. "Set" will overwrite, add will create if not
        already defined.
     */
    httpAddHeaderString(conn, "Content-Type", "text/html");
    httpSetHeaderString(conn, "Cache-Control", "no-cache");

    httpWrite(q, "<html><title>v1/action</title><body>\r\n");
    httpWrite(q, "<p>param: %s</p>\n",
              httpGetParam(conn, "param", "undefined"));
    httpWrite(q, "<p>method: %s</p>\n", conn->rx->method);
    httpWrite(q, "<p>client: %s:%d</p>\n", conn->ip, conn->port);
    httpWrite(q, "</body></html>\r\n");

    /*
        Call finalize output and close the request.
        Delay closing if you want to do asynchronous output and close later.
     */
    httpFinalize(conn);

#if 0
    /*
        Useful things to do in actions
     */
    httpRedirect(conn, 302, "/other-uri");
    httpError(conn, 409, "My message : %d", 5);
#endif
}

void test(HttpConn *conn)
{
    httpSetStatus(conn, 200);
    httpAddHeaderString(conn, "Content-Type", "text/plain");
    httpWrite(conn->writeq, "Hello World\n");
    httpFinalize(conn);
}
/*
    Create a simple stand-alone web server
 */
int main(int argc, char **argv)
{
    Mpr *mpr;
    int rc;

    rc = MPR_ERR_CANT_CREATE;
    if ((mpr = mprCreate(0, NULL, MPR_USER_EVENTS_THREAD)) == 0)
    {
        mprError("Cannot create runtime");
        return -1;
    }
    if (httpCreate(HTTP_CLIENT_SIDE | HTTP_SERVER_SIDE) == 0)
    {
        mprError("Cannot create the HTTP services");
        return -1;
    }
    mprStart();

    if (maParseConfig("mo.conf") < 0)
    {
        mprError("Cannot parse the config file %s", "mo.conf");
        return -1;
    }

    httpDefineAction("/v1/action", doAction);
    httpDefineAction("/v1/test", test);


    if (httpStartEndpoints() < 0)
    {
        mprError("Cannot start the web server");
        return -1;
    }
    mprServiceEvents(-1, 0);
    mprDestroy();
    return 0;
}
