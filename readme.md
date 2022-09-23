## Go Away presentation

Go Away is a conditional link redirection service.
Concretely, it's possible for a given route to specify a redirection address as on a link shortener, with the difference here that it's possible to redirect to a link according to the user's context.

For example, imagine a QR Code available on our restaurant table. It contains the link `https://goaway.unresto.fr/menu`.
The `/menu` route redirects to the French menu by default. However, the user is redirected to the English menu if the request contains an Accept-Language header not equal to fr-FR.

