package email

import "fmt"

// GetRegisterHtmlBody returns the registration email body with the provided name and username.
func GetRegisterHtmlBody(name, username string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Welcome to Speakeasy</title>
  </head>
  <body style="font-family: Arial, sans-serif; color: white; margin: 0; padding: 0;">
    <table role="presentation" width="100%%" cellspacing="0" cellpadding="0" border="0">
      <tr>
        <td style="padding: 20px;">
          <table role="presentation" width="600" cellspacing="0" cellpadding="0" border="0" align="center" style="border: 1px solid #ddd; border-radius: 5px; background-image: url('https://pixelsharing.wordpress.com/wp-content/uploads/2010/11/gothic-background-tileable-and-seamless-pattern.jpg'); background-repeat: repeat;">
            <tr>
              <td align="center" style="position: relative;">
                <img src="https://cdn1.iconfinder.com/data/icons/wild-west-24/64/hat-western-cowboy-512.png" alt="Speakeasy Logo" style="position: absolute; top: 50%%; left: 50%%; transform: translate(-50%%, -50%%); width: 40%%; height: auto; z-index: 0; opacity: 0.3;" />
                <div style="position: relative; z-index: 1; padding: 20px;">
                  <h1 style="text-align: center;">Welcome to Speakeasy, %s!</h1>
                  <p style="text-align: center;">Thank you for creating an account with Speakeasy!</p>
                  <p style="text-align: center;">Your account '%s' has just been registered and you're ready to go.</p>
                  <p style="text-align: center;">Just navigate back to <a href="https://localhost:5173/" style="color: #ffffff;">Speakeasy</a> to get started.</p>
                </div>
              </td>
            </tr>
            <tr>
              <td align="center" style="padding: 20px; font-size: 12px; color: #888;">
                <p>&copy; 2024 Speakeasy. All rights reserved.</p>
              </td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
  </body>
</html>
`, name, username)
}
