��    g      T  �   �      �     �     �  )   �  #   	  1   '	  3   Y	     �	  
   �	    �	     7  i   J  �   �  %   �  /   �     �     �      A        T  H   i  9   �  ,   �  +        E     `  8   u     �  8   �     �  3     +   A  7   m     �     �     �     �     �       
   #  6   .  $  e  "   �     �     �      �  >   �  &   5  7   \     �     �  $   �     �  ,     V  1  1   �     �  +   �     �  2     +   E  #   q  w   �       %   '  )   M  .   w     �  ?   �  �   �  %   �  $     '   C  ,   k     �     �  d   �  B   1  ?   t  3   �  g   �  '   P     x  *   �     �  	   �     �     �                       .      =      V      g   �   t   %   �   $   !!     F!  #   ]!  "   �!     �!     �!  �  �!  +   T#     �#  -   �#  &   �#  9   �#  7   $     O$     g$  �  u$     �&  T   '  �   c'  /   I(  '   y(     �(  1   �(  8  �(  Q   &*     x*  W   �*  <   �*  /   #+  D   S+  "   �+      �+  7   �+     ,  5   -,     c,  ;   r,  5   �,  1   �,     -     /-     J-  !   ^-     �-     �-     �-  4   �-  [  �-  -   S2     �2     �2  $   �2  F   �2  +   3  A   B3  &   �3     �3  2   �3     �3  -   4  4  D4  +   y5     �5  -   �5     �5      6  3    6  0   T6  �   �6     7  8   )7  D   b7  9   �7     �7  C   �7    98  !   H9  #   j9  '   �9  3   �9     �9      :  q   :  8   �:  C   �:  2   ;  m   ;;  /   �;  %   �;  8   �;     8<     V<     g<     y<  
   �<     �<     �<     �<     �<     �<     
=  �   =  2   �=  1   �=  '   >  1   9>  /   k>     �>     �>     X   &          >   *   4   V   \       0   S                 2   
   ,   E          F   K   L   	              W      b   8           O             M           3   Y   d          ;   a      6       A   e      D   f      `   =       (   $   _   J       ?                !   9   Q   1       I   U         C           <           -   Z   5           B   )      T   [      g   :       '      ]   /   %              ^   .      "         H   +   c       R              @          P              G   N   #           7       %s	%s	%s (forks not shown: %d)	 %s: %s
 '%s' is no longer allowed to access '%s'
 '%s' is now allowed to access '%s'
 '%s' previously allowed access to '%s'. Skipping
 '%s:' is not allowed to access additional hardware
 (deprecated) please use "list" 2fa code:  A concise summary of key attributes of the snappy system, such as the release and channel.

The verbose output includes the specific version information for the factory image, the running image and the image that will be run on reboot, together with a list of the available channels for this image.

Providing a package name will display information about a specific installed package.

The verbose version of the info command for a package will also tell you the available channels for that package, when it was installed for the first time, disk space utilization, and in the case of frameworks, which apps are able to use the framework. Activate a package Activate a package that has previously been deactivated. If the package is already activated, do nothing. Allows rollback of a snap to a previous installed version. Without any arguments, the previous installed version is selected. It is also possible to specify the version to rollback to as a additional argument.
 Assign a hardware device to a package Assign hardware to a specific installed package Builds a snap package Can't read hook file %q: %v Configures a package. The configuration is a YAML file, provided in the specified file which can be "-" for stdin. Output of the command is the current configuration, so running this command with no input file provides a snapshot of the app's current config. Creates a snap package and if available, runs the review scripts. Deactivate a package Deactivate a package. If the package is already deactivated, do nothing. Display a summary of key attributes of the snappy system. Do not clean up old versions of the package. Ensures system is running with latest parts First boot has already run Generated '%s' snap
 Include information about packages from the snappy store Install a snap package Install snaps even if the signature can not be verified. Installing %s
 List active components installed on a snappy system List assigned hardware device for a package List assigned hardware for a specific installed package Log into the store Login successful Name	Date	Version	 Name	Date	Version	Developer	 Name	Version	Summary	 No snap: '%s' found Password:  Provide information about a specific installed package Provides a list of all active components installed on a snappy system.

If requested, the command will find out if there are updates for any of the components and indicate that by appending a * to the date. This will be slower as it requires a round trip to the app store on the network.

The developer information refers to non-mainline versions of a package (much like PPAs in deb-based Ubuntu). If the package is the primary version of that package in Ubuntu then the developer info is not shown. This allows one to identify packages which have custom, non-standard versions installed. As a special case, the “sideload” developer refers to packages installed manually on the system.

When a verbose listing is requested, information about the channel used is displayed; which is one of alpha, beta, rc or stable, and all fields are fully expanded too. In some cases, older (inactive) versions of snappy packages will be installed, these will be shown in the verbose output and the active version indicated with a * appended to the name of the component. Provides more detailed information Purge an installed package. Purging %s
 Query and modify snappy services Query and modify snappy services of locally-installed packages Query the store for available packages Reboot if necessary to be on the latest running system. Reboot to use %s version %s. Reboot to use the new %s. Rebooting to satisfy updates for %s
 Remove a snapp part Remove all the data from the listed packages Remove all the data from the listed packages. Normally this is used for packages that have been removed and attempting to purge data for an installed package will result in an error. The --installed option  overrides that and enables the administrator to purge all data for an installed package (effectively resetting the package completely). Remove hardware from a specific installed package Removing %s
 Rollback to a previous version of a package Search for packages to install Set configuration for a specific installed package Set configuration for an installed package. Set properties of system or package Set properties of system or package

Supported properties are:
  active=VERSION

Example:
  set hello-world active=1.0
 Setting %s to version %s
 Show all available forks of a package Show available updates (requires network) Show channel information and expand all fields Snap	Service	State Specify an alternate output directory for the resulting package The "versions" command is no longer available.

Please use the "list" command instead to see what is installed.
The "list -u" (or "list --updates") will show you the available updates
and "list -v" (or "list --verbose") will show all installed versions.
 The Package to install (name or path) The configuration for the given file The configuration for the given install The hardware device path (e.g. /dev/ttyUSB0) The package to rollback  The version to rollback to This command adds access to a specific hardware device (e.g. /dev/ttyUSB0) for an installed package. This command is no longer available, please use the "list" command This command list what hardware an installed package can access This command logs the given username into the store This command removes access of a specific hardware device (e.g. /dev/ttyUSB0) for an installed package. Unassign a hardware device to a package Update all installed parts Use --show-all to see all available forks. Username for the login apps: %s
 architecture: %s
 binary-size: %v
 channel: %s
 data-size: %s
 frameworks: %s
 installed: %s
 package name is required produces manpage release: %s
 snappy autopilot triggered a reboot to boot into an up to date system -- temprorarily disable the reboot by running 'sudo shutdown -c' unable to disable %s's service %s: %v unable to enable %s's service %s: %v unable to get logs: %v unable to start %s's service %s: %v unable to stop %s's service %s: %v updated: %s
 version: %s
 Project-Id-Version: snappy
Report-Msgid-Bugs-To: FULL NAME <EMAIL@ADDRESS>
POT-Creation-Date: 2015-10-15 15:53+0200
PO-Revision-Date: 2015-10-21 16:12+0000
Last-Translator: Marcos Lans <Unknown>
Language-Team: Galician <gl@li.org>
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 8bit
X-Launchpad-Export-Date: 2015-10-22 05:57+0000
X-Generator: Launchpad (build 17812)
 %s	%s	%s (bifurcacións non mostradas: %d)	 %s: %s
 «%s» xa non ten permiso de acceso a «%s»
 «%s» ten permiso de acceso a «%s»
 «%s» permitiu o acceso previamente a «%s ». Saltando
 «%s:» non ten permiso de acceso a hardware adicional
 (obsoleto) use «list» Código 2fa:  Resumo conciso dos atributos fundamentais do sistema snappy, como a publicación e a canle.

A saída con información detallada inclúe a información específica da versión para a imaxe de fábrica, a imaxe en execución e a imaxe que se executará ao reiniciar, xunto cunha lista de canles para esta imaxe.

Indicar un nome de paquete mostrará información sobre o paquete instalado específico.

A versión con información detallada da orde para o paquete tamén mostrará as canles dispoñíbeis dese paquete, cando se instalou por primeira vez, a utilización de espazo no disco e no caso de frameworks, que aplicativos usa o framework. Activar un paquete Activar un paquete previamente desactivado. Se o paquete xa está activado, ignorar. Permite a reversión dun snap a unha versión instalada anteriormente. Se non se usan argumentos seleccionarase a versión instalada previamente. Tamén se pode especificar a versión á que reverter como un argumento adicional.
 Asignar un dispositivo de hardware a un paquete Asignar hardware a un paquete instalado Constrúe un paquete snap Non é posíbel ler o ficheiro do «hook» %q: %v Configura un paquete.  A configuración é un ficheiro YAML, proporcionado no ficheiro específico que pode ser «-» para stdin. A saída da orde é a configuración actual, de xeito que se se executa esta orde sen indicar un ficheiro de entrada proporcionará unha imaxe da configuración actual do aplicativo. Crea un paquete snap e, se están dispoñíbeis, executa os scripts de revisión. Desactivar un paquete Desactivar un paquete previamente activado. Se o paquete xa está desactivado, ignorar. Mostra un resumo dos atributos principais do sistema snappy. Non eliminar as versións antigas dos paquetes. Asegúrase de que o sistema está en execución coa últimas partes. Xa se executou o primeiro arranque Xerado o paquete snap de «%s»
 Incluír información sobre paquetes da tenda de snappy Instalar un paquete snap Instalar snaps incluso con sinaturas non comprobadas. Instalando %s
 Lista os compoñentes activos instalados nun sistema snappy Lista o dispositivo de hardware asignado a un paquete Listar o hardware asignado a un paquete instalado Iniciar sesión na tenda Inicio de sesión correcto Nome	Data	Versión	 Nome	Data	Versión	Desenvolvedor	 Nome	Versión	Resumo	 Ningún snap: atopouse «%s» Contrasinal:  Proporcionar información sobre un paquete instalado Proporciona unha lista dos compoñentes activos instalados nun sistema snappy.

Se se solicita, unha orde averiguará se hai actualizacións para algún dos compoñentes e indicarao anexando un * á data. Isto pode ser lento xa que require un proceso de ida e volta á tenda de aplicativos na rede.

A información do desenvolvedor refírese a versións que non son da liña principal dun paquete (como as PPAs de Ubuntu baseado en Debian). Se o paquete é unha primeira versión do paquete en Ubuntu non se mostrará a información do desenvolvedor. Isto permite identificar paquetes con versións personalizadas non estándar instaladas. Como un caso especial, o desenvolvedor «sideload» refírese a paquetes instalados manualmente no sistema.

Cando se solicite unha lista detallada, mostrarase a información sobre a canle utilizada; que pode ser alfa, beta, rc ou estábel e todos os campos estarán expandidos. Nalgúns casos, instalaranse versións antigas (inactivas) dos paquetes snappy; mostraranse na saída de información detallada e a versión activa indicarase cun * anexado ao nome do compoñente. Proporciona unha información máis detallada Purga un paquete instalado. Purgando %s
 Consultar e modificar sevizos snappy Consultar e modificar sevizos snappy de paquetes instalados localmente Consulta na tenda os paquetes dispoñíbeis Reiniciar se é necesario estar no último sistema en execución. Reiniciar para usar %s na versión %s. Reiniciar para usar o novo %s. Reiniciando para rematar as actualizacións de %s
 Retirar unha parte snapp Eliminar todos os datos dos paquetes da lista Eliminar todos os datos dos paquetes da lista. Normalmente úsase para paquetes eliminados que cando se tentan purgar os seus datos obtemos un erro. A opción --installed anula iso e activa que o aministrador poida purgar todos os datos dun paquete instalado (restabelecendo eficaz e completamente o paquete) Eliminar hardware para un paquete instalado Eliminando %s
 Reverter a unha versión anterior dun paquete Buscar paquetes para instalar Configurar un paquete instalado Estabelecer a configuración dun paquete instalado. Estabelecer as propiedades do sistema ou paquete Estabelece as propiedade do sistema ou paquete

As propiedades aceptadas son:
  active=VERSION

Exemplo:
  set hello-world active=1.0
 Axustando %s á versión %s
 Mostrar todas as bifurcacións dispoñíbeis dun paquete Mostrar as actualizacións dispoñíbeis (precisa conexión á rede) Mostrar a información da canle e ampliar todos os campos Snap	Servizo	Estado Indicar un cartafol de saída alternativo para o paquete resultante A orde «versions» xa non está dispoñíbel.

Use a orde «list» no seu canto para ver cal está instalada.
«list -u» (ou «list --updates») mostrará as actualizacións dispoñíbeis.
e «list -v» (ou «list --verbose») mostrará todas as versións instaladas.
 Paquete a instalar (nome ou ruta) Configuración para o ficheiro dado Configuración para a instalación dada Ruta do dispositivo de hardware (p.e. /dev/ttyUSB0) Paquete que reverter  Versión á que reverter Esta orde engade o acceso a un dispositivo de hardware específico para un paquete instalado (p.e. /dev/ttyUSB0). Esta orde xa non está dispoñíbel, use a orde «list» Esta orde lista o hardware ao que pode acceder un paquete instalado Esta orde rexistra o nome de usuario dado na tenda Esta orde elimina o acceso a un dispositivo específico de hardware (p.e /dev/ttyUSB0) dun paquete instalado. Desligar un dispositivo de hardware dun paquete Actualizar todas as partes instaladas Usar --show-all para ver as bifurcacións dispoñíbeis. Nome de usuario para o acceso aplicativos: %s
 arquitectura: %s
 tamaño do binario: %v
 canle: %s
 data-tamaño: %s
 contornos de traballo: %s
 instalado: %s
 precísase o nome do paquete produce unha manpage publicación: %s
 snappy autopilot activou un reinicio para comezar nun sistema actualizado -- desactive temporalmente o reinicio executando «sudo shutdown -c» Non foi posíbel desactivar o servizo de %s %s: %v non foi posíbel activar o servizo para %s %s: %v non foi posíbel obter os rexistros: %v non foi posíbel iniciar o servizo para %s %s: %v non foi posíbel parar o servizo para %s %s: %v actualizado: %s
 versión: %s
 