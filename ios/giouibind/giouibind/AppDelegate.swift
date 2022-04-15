//
//  AppDelegate.swift
//  giouibind
//
//  Created by Jessica Fleming on 10.04.22.
//

import UIKit
import Mobile

@main
class AppDelegate: UIResponder, UIApplicationDelegate {

    var window:UIWindow?

    func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {
        
        self.window = UIWindow(frame: CGRect(
            x: UIScreen.main.bounds.minX,
            y: UIScreen.main.bounds.minY + 40,
            width: UIScreen.main.bounds.width,
            height: UIScreen.main.bounds.height - 40
        ))
        window?.rootViewController = ViewController()
        window?.makeKeyAndVisible()
        MobileInventGod(GodObject())

        return true
        }
    
    var gioView : ViewController {
        get {
            return self.window!.rootViewController! as! ViewController
        }
    }
}

