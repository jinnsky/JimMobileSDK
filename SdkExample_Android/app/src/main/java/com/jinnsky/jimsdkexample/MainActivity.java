package com.jinnsky.jimsdkexample;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.Toast;

import go.jimsdk.*;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Button emailVerificationButton = (Button) findViewById(R.id.emailVerificationButton);

        Client client = null;

        try {
            client = Jimsdk.newClient("http://api2.jimyun.com",
                                             23,
                                             "iu3TKjwRUCGfIwtTH9gXeYsq",
                                             "kJek81coyFG4V3eSg79b82HU");
        } catch (Exception e) {
            e.printStackTrace();
        }

        final Client finalClient = client;
        emailVerificationButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                finalClient.sendVerifyEmailAsync("yangjingtian@oudmon.com", new VerifyEmailResponseListener() {
                    @Override
                    public void onFailure(String s) {
                        Toast.makeText(getApplicationContext(), "Sent verification email - Failed. " + s, Toast.LENGTH_LONG).show();
                    }

                    @Override
                    public void onSuccess(VerifyEmailResponseData responseData) {
                        if (responseData.getResult()) {
                            Toast.makeText(getApplicationContext(), "Sent verification email - OK.", Toast.LENGTH_LONG).show();
                        } else {
                            Toast.makeText(getApplicationContext(), "Sent verification email - Failed.", Toast.LENGTH_LONG).show();
                        }
                    }
                });
            }
        });
    }
}
