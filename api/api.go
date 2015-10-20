/* Exposes the Simple Clusterer and Stream Clusterer APIs */
package api;

import (
    "github.com/wenkesj/rphash/simple"
    "github.com/wenkesj/rphash/stream"
);

func Stream() *stream.Stream {
    return stream.NewStream();
};

func Simple() *simple.Simple {
    return simple.NewSimple();
};
