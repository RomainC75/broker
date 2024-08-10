import { AuthenticationResult, Configuration, EventType, PublicClientApplication } from "@azure/msal-browser";
import { MsalProvider } from "@azure/msal-react";
import { ReactNode } from "react";

const msalConfig: Configuration = {
    auth: {
        clientId: import.meta.env.VITE_API_SSO_AUDIENCE_ID ?? "",
        authority: `https://login.microsoftonline.com/${import.meta.env.VITE_API_SSO_TENANT_ID ?? ""}`,
        redirectUri: "/",
        postLogoutRedirectUri: "/"
    },
    system: {
        allowNativeBroker: false 
    }
};

interface IAuthProvider {
    children: ReactNode
}

const publicClientApplication = new PublicClientApplication(msalConfig)

const AuthProvider = ({children}: IAuthProvider)  => {
    publicClientApplication.enableAccountStorageEvents();
    publicClientApplication.addEventCallback((event) => {
        const authenticationResult = event.payload as AuthenticationResult;
        const account = authenticationResult?.account;
        // setUser(account)
        console.log("-> account !!", authenticationResult)
        localStorage.setItem("token", JSON.stringify(authenticationResult));
        if (event.eventType === EventType.LOGIN_SUCCESS && account) {
            publicClientApplication.setActiveAccount(account);
        }
    });

    return(
        <MsalProvider instance={publicClientApplication}>{children}</MsalProvider>
    )
}

export default AuthProvider